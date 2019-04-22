package server

import (
	"context"

	pb "github.com/nilslice/schema-registry/v1/go/registrypb"
)

func (s Service) Commit(context.Context, *pb.CommitRequest) (*pb.CommitResponse, error) {
	return nil, nil
}
