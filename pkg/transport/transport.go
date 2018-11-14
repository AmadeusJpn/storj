// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package transport

import (
	"context"
	"time"

	"github.com/zeebo/errs"
	"google.golang.org/grpc"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/provider"
)

var (
	mon = monkit.Package()
	//Error is the errs class of standard Transport Client errors
	Error = errs.Class("transport error")
)

const dialTimeout = 5 * time.Second

// Client defines the interface to an transport client.
type Client interface {
	DialNode(ctx context.Context, node *pb.Node, opts ...grpc.DialOption) (*grpc.ClientConn, error)
	DialAddress(ctx context.Context, address string, opts ...grpc.DialOption) (*grpc.ClientConn, error)
	Identity() *provider.FullIdentity
}

// Transport interface structure
type Transport struct {
	identity *provider.FullIdentity
}

// NewClient returns a newly instantiated Transport Client
func NewClient(identity *provider.FullIdentity) Client {
	return &Transport{identity: identity}
}

// DialNode returns a grpc connection with tls to a node
func (o *Transport) DialNode(ctx context.Context, node *pb.Node, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	defer mon.Task()(&ctx)(&err)

	if node.Address == nil || node.Address.Address == "" {
		return nil, Error.New("no address")
	}

	// add ID of node we are wanting to connect to
	dialOpt, err := o.identity.DialOption(node.GetId())
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(ctx, node.GetAddress().Address, append([]grpc.DialOption{dialOpt, grpc.WithTimeout(dialTimeout)}, opts...)...)
}

// DialAddress returns a grpc connection with tls to an IP address
func (o *Transport) DialAddress(ctx context.Context, address string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	defer mon.Task()(&ctx)(&err)

	dialOpt, err := o.identity.DialOption("")
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(ctx, address, append([]grpc.DialOption{dialOpt, grpc.WithTimeout(dialTimeout)}, opts...)...)
}

// Identity is a getter for the transport's identity
func (o *Transport) Identity() *provider.FullIdentity {
	return o.identity
}

// Close implements io.closer, closing the transport connection(s)
