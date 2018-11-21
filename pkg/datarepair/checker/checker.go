// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package checker

import (
	"context"
	"time"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"

	"storj.io/storj/pkg/auth"
	"storj.io/storj/pkg/datarepair/queue"
	"storj.io/storj/pkg/dht"
	"storj.io/storj/pkg/irreparabledb"
	"storj.io/storj/pkg/node"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb"
	"storj.io/storj/pkg/utils"
	"storj.io/storj/storage"
)

// Checker is the interface for data repair checker
type Checker interface {
	Run(ctx context.Context) error
}

// Checker contains the information needed to do checks for missing pieces
type checker struct {
	pointerdb   *pointerdb.Server
	repairQueue *queue.Queue
	overlay     pb.OverlayServer
	irrdb       *irreparabledb.Server
	limit       int
	logger      *zap.Logger
	ticker      *time.Ticker
}

// newChecker creates a new instance of checker
func newChecker(pointerdb *pointerdb.Server, repairQueue *queue.Queue, overlay pb.OverlayServer, irrdb *irreparabledb.Server, limit int, logger *zap.Logger, interval time.Duration) *checker {
	return &checker{
		pointerdb:   pointerdb,
		repairQueue: repairQueue,
		overlay:     overlay,
		irrdb:       irrdb,
		limit:       limit,
		logger:      logger,
		ticker:      time.NewTicker(interval),
	}
}

// Run the checker loop
func (c *checker) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	for {
		err = c.identifyInjuredSegments(ctx)
		if err != nil {
			zap.L().Error("Checker failed", zap.Error(err))
		}

		select {
		case <-c.ticker.C: // wait for the next interval to happen
		case <-ctx.Done(): // or the checker is canceled via context
			return ctx.Err()
		}
	}
}

// identifyInjuredSegments checks for missing pieces off of the pointerdb and overlay cache
func (c *checker) identifyInjuredSegments(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)
	c.logger.Debug("entering pointerdb iterate")

	err = c.pointerdb.Iterate(ctx, &pb.IterateRequest{Recurse: true},
		func(it storage.Iterator) error {
			var item storage.ListItem
			lim := c.limit
			if lim <= 0 || lim > storage.LookupLimit {
				lim = storage.LookupLimit
			}
			for ; lim > 0 && it.Next(&item); lim-- {
				pointer := &pb.Pointer{}
				err = proto.Unmarshal(item.Value, pointer)
				if err != nil {
					return Error.New("error unmarshalling pointer %s", err)
				}
				remote := pointer.GetRemote()
				if remote == nil {
					c.logger.Debug("no remote segment on pointer")
					continue
				}
				pieces := remote.GetRemotePieces()
				if pieces == nil {
					c.logger.Debug("no pieces on remote segment")
					continue
				}
				var nodeIDs []dht.NodeID
				for _, p := range pieces {
					nodeIDs = append(nodeIDs, node.IDFromString(p.NodeId))
				}
				missingPieces, err := c.offlineNodes(ctx, nodeIDs)
				if err != nil {
					return Error.New("error getting offline nodes %s", err)
				}
				numHealthy := len(nodeIDs) - len(missingPieces)
				if (int32(numHealthy) >= pointer.Remote.Redundancy.MinReq) && (int32(numHealthy) < pointer.Remote.Redundancy.RepairThreshold) {
					err = c.repairQueue.Enqueue(&pb.InjuredSegment{
						Path:       string(item.Key),
						LostPieces: missingPieces,
					})
					if err != nil {
						return Error.New("error adding injured segment to queue %s", err)
					}
				} else if int32(numHealthy) < pointer.Remote.Redundancy.MinReq {
					// make an entry in to the irreparabledb table
					rmtSegInfo := &pb.RmtSegInfo{
						Key:                item.Key,
						Val:                item.Value,
						LostPiecesCount:    int64(len(missingPieces)),
						RepairUnixSec:      time.Now().Unix(),
						RepairAttemptCount: int64(1),
					}

					apiKey, ok := auth.GetAPIKey(ctx)
					if !ok {
						return Error.New("no api key was provided")
					}

					putReq := &pb.PutIrrSegRequest{
						Info:   rmtSegInfo,
						APIKey: apiKey,
					}
					_, err = c.irrdb.Put(ctx, putReq)
					if err != nil {
						return Error.New("couldn't make an entry into irreparable db: %s", err)
					}
				}
			}
			return nil
		},
	)
	return err
}

// returns the indices of offline nodes
func (c *checker) offlineNodes(ctx context.Context, nodeIDs []dht.NodeID) (offline []int32, err error) {
	responses, err := c.overlay.BulkLookup(ctx, utils.NodeIDsToLookupRequests(nodeIDs))
	if err != nil {
		return []int32{}, err
	}
	nodes := utils.LookupResponsesToNodes(responses)
	for i, n := range nodes {
		if n == nil {
			offline = append(offline, int32(i))
		}
	}
	return offline, nil
}
