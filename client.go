package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/weeon/contract"
	"google.golang.org/grpc"
)

var (
	ServiceNameNotFound = errors.New("Service Name Not Found ")
	grpcAddrsKeyFormat  = "%s_grpc_addrs"
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

func NewClientManager(ctx context.Context, namespace string, config contract.Config) (*ClientManager, error) {
	b, err := config.Get(fmt.Sprintf(grpcAddrsKeyFormat, namespace))
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
