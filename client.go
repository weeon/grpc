package grpc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/weeon/contract"
	"google.golang.org/grpc"
)

var (
	ServiceNameNotFound = errors.New("Service Name Not Found ")
	grpcAddrsKey        = "grpc_addrs"
)

type ClientManager struct {
	config contract.Config

	addrs map[string]string
}

func (c *ClientManager) GetGrpcConn(ctx context.Context, name string) (*grpc.ClientConn, error) {
	addr, ok := c.addrs[name]
	if !ok {
		return nil, ServiceNameNotFound
	}
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	return conn, err
}

func NewClientManager(ctx context.Context, config contract.Config) (*ClientManager, error) {
	b, err := config.Get(grpcAddrsKey)
	if err != nil {
		return nil, err
	}
	addrs := make(map[string]string)
	err = json.Unmarshal(b, &addrs)
	if err != nil {
		return nil, err
	}
	return &ClientManager{
		config: config,
		addrs:  addrs,
	}, nil
}
