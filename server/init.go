package server

import (
	"context"

	pb "github.com/nilslice/schema-registry/v1/go/registrypb"
)

func (s Service) Init(context.Context, *pb.InitRequest) (*pb.InitResponse, error) {
	return nil, nil
}
