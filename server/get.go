package server

import (
	"context"

	pb "github.com/nilslice/schema-registry/go/registrypb"
)

func (s Service) Get(context.Context, *pb.GetRequest) (*pb.GetResponse, error) {
	return nil, nil
}
