package server

import (
	"bytes"
	"context"

	"github.com/nilslice/protolock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/nilslice/schema-registry/v1/go/registrypb"
)

func (s Service) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	if req.GetSchema() == nil {
		return nil, status.Error(codes.NotFound, "no schema in request")
	}
	current := bytes.NewReader(req.GetLockfile())
	updated := bytes.NewReader(req.GetSchema().GetLockfile())

	cur, err := protolock.FromReader(current)
	if err != nil {
		return nil, status.Error(codes.Internal, "current "+err.Error())
	}
	upd, err := protolock.FromReader(updated)
	if err != nil {
		return nil, status.Error(codes.Internal, "updated "+err.Error())
	}

	report, err := protolock.Compare(cur, upd)
	if err != nil && err != protolock.ErrWarningsFound {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var warnings []*pb.Warning
	for _, warning := range report.Warnings {
		warnings = append(warnings, &pb.Warning{
			Filepath: protolock.OSPath(warning.Filepath).String(),
			Message:  warning.Message,
		})
	}

	return &pb.StatusResponse{
		Warnings: warnings,
	}, nil
}
