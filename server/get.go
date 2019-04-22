package server

import (
	"context"

	pb "github.com/nilslice/schema-registry/v1/go/registrypb"
)

func (s Service) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, nil
}
