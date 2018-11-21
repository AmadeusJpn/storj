// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package irreparabledb

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/internal/migrate"
	"storj.io/storj/pkg/auth"
	dbx "storj.io/storj/pkg/irreparabledb/dbx"
	"storj.io/storj/pkg/pb"
	pdbauth "storj.io/storj/pkg/pointerdb/auth"
)

var (
	mon = monkit.Package()
)

// Server implements the statdb RPC service
type Server struct {
	DB     *dbx.DB
	logger *zap.Logger
}

// NewServer creates instance of Server
func NewServer(driver, source string, logger *zap.Logger) (*Server, error) {
	db, err := dbx.Open(driver, source)
	if err != nil {
		return nil, err
	}

	err = migrate.Create("irreparabledb", db)
	if err != nil {
		return nil, err
	}

	return &Server{
		DB:     db,
		logger: logger,
	}, nil
}

func (s *Server) validateAuth(ctx context.Context) error {
	APIKeyBytes, ok := auth.GetAPIKey(ctx)
	if !ok {
		return Error.New("no api key was provided")
	}
	if !pdbauth.ValidateAPIKey(string(APIKeyBytes)) {
		s.logger.Error("unauthorized request: ", zap.Error(status.Errorf(codes.Unauthenticated, "Invalid API credential")))
		return status.Errorf(codes.Unauthenticated, "Invalid API credential")
	}
	return nil
}

// Put a db entry for the provided remote segment info
func (s *Server) Put(ctx context.Context, putReq *pb.PutIrrSegRequest) (resp *pb.PutIrrSegResponse, err error) {
	info := putReq.Info

	getReq := &pb.GetIrrSegRequest{
		Key:    info.Key,
		APIKey: putReq.APIKey,
	}

	// entry exists already?
	getResp, _ := s.Get(ctx, getReq)
	if getResp != nil {
		rmtSegInfo := putReq.Info
		rmtSegInfo.RepairAttemptCount = rmtSegInfo.GetRepairAttemptCount() + int64(1)
		updateReq := &pb.PutIrrSegRequest{
			Info:   rmtSegInfo,
			APIKey: putReq.APIKey,
		}
		return s.Update(ctx, updateReq)
	}
	return s.Create(ctx, putReq)
}

// Create a db entry for the provided remote segment info
func (s *Server) Create(ctx context.Context, putReq *pb.PutIrrSegRequest) (resp *pb.PutIrrSegResponse, err error) {
	defer mon.Task()(&ctx)(&err)
	s.logger.Debug("entering irreparabledb Create")

	if err := s.validateAuth(ctx); err != nil {
		return nil, err
	}

	info := putReq.Info
	_, err = s.DB.Create_Irreparabledb(
		ctx,
		dbx.Irreparabledb_Segmentkey(info.Key),
		dbx.Irreparabledb_Segmentval(info.Val),
		dbx.Irreparabledb_PiecesLostCount(info.LostPiecesCount),
		dbx.Irreparabledb_SegDamagedUnixSec(info.RepairUnixSec),
		dbx.Irreparabledb_SegCreatedAt(time.Unix(info.RepairUnixSec, 0)),
		dbx.Irreparabledb_RepairAttemptCount(info.RepairAttemptCount),
	)
	if err != nil {
		return &pb.PutIrrSegResponse{
			Status: pb.PutIrrSegResponse_FAIL,
		}, status.Errorf(codes.Internal, err.Error())
	}

	s.logger.Debug("created in the db: " + string(info.Key))
	return &pb.PutIrrSegResponse{
		Status: pb.PutIrrSegResponse_OK,
	}, nil
}

// Get a irreparable's segment info from the db
func (s *Server) Get(ctx context.Context, getReq *pb.GetIrrSegRequest) (resp *pb.GetIrrSegReponse, err error) {
	defer mon.Task()(&ctx)(&err)
	s.logger.Debug("entering irreparabaledb Get")

	err = s.validateAuth(ctx)
	if err != nil {
		return nil, err
	}

	dbSegInfo, err := s.DB.Get_Irreparabledb_By_Segmentkey(ctx, dbx.Irreparabledb_Segmentkey(getReq.GetKey()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	rmtseginfo := &pb.RmtSegInfo{
		Key:                dbSegInfo.Segmentkey,
		Val:                dbSegInfo.Segmentval,
		LostPiecesCount:    dbSegInfo.PiecesLostCount,
		RepairUnixSec:      dbSegInfo.SegDamagedUnixSec,
		RepairAttemptCount: dbSegInfo.RepairAttemptCount,
	}
	return &pb.GetIrrSegReponse{
		Info:   rmtseginfo,
		Status: pb.GetIrrSegReponse_OK,
	}, nil
}

// Delete a irreparable's segment info from the db
func (s *Server) Delete(ctx context.Context, delReq *pb.DeleteIrrSegRequest) (resp *pb.DeleteIrrSegResponse, err error) {
	defer mon.Task()(&ctx)(&err)
	s.logger.Debug("entering irreparabaledb Delete")

	err = s.validateAuth(ctx)
	if err != nil {
		return nil, err
	}

	_, err = s.DB.Delete_Irreparabledb_By_Segmentkey(ctx, dbx.Irreparabledb_Segmentkey(delReq.GetKey()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.DeleteIrrSegResponse{
		Status: pb.DeleteIrrSegResponse_OK,
	}, nil
}

// Update a db entry for the provided remote segment info
func (s *Server) Update(ctx context.Context, putReq *pb.PutIrrSegRequest) (resp *pb.PutIrrSegResponse, err error) {
	defer mon.Task()(&ctx)(&err)
	s.logger.Debug("entering irreparabledb Update")

	if err := s.validateAuth(ctx); err != nil {
		return nil, err
	}

	info := putReq.Info
	update := dbx.Irreparabledb_Update_Fields{
		Segmentval:         dbx.Irreparabledb_Segmentval(info.Val),
		PiecesLostCount:    dbx.Irreparabledb_PiecesLostCount(info.LostPiecesCount),
		SegDamagedUnixSec:  dbx.Irreparabledb_SegDamagedUnixSec(info.RepairUnixSec),
		SegCreatedAt:       dbx.Irreparabledb_SegCreatedAt(time.Unix(info.RepairUnixSec, 0)),
		RepairAttemptCount: dbx.Irreparabledb_RepairAttemptCount(info.RepairAttemptCount),
	}
	_, err = s.DB.Update_Irreparabledb_By_Segmentkey(
		ctx,
		dbx.Irreparabledb_Segmentkey(info.Key),
		update,
	)
	if err != nil {
		return &pb.PutIrrSegResponse{
			Status: pb.PutIrrSegResponse_FAIL,
		}, status.Errorf(codes.Internal, err.Error())
	}

	s.logger.Debug("created in the db: " + string(info.Key))
	return &pb.PutIrrSegResponse{
		Status: pb.PutIrrSegResponse_OK,
	}, nil
}
