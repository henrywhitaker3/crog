package server

import (
	"context"

	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/pb"
)

func (s Server) List(context.Context, *pb.ListActionsRequest) (*pb.ListActionsResponse, error) {
	acts := []*pb.Action{}

	for _, a := range s.cfg.Actions {
		acts = append(acts, &pb.Action{
			Name:    a.Name,
			Command: a.Command,
			Code:    int64(a.Code),
		})
	}

	return &pb.ListActionsResponse{Actions: acts}, nil
}

func (s Server) Run(ctx context.Context, req *pb.RunActionRequest) (*pb.RunActionResponse, error) {
	action, err := s.cfg.GetAction(req.Action)
	if err != nil {
		return nil, err
	}
	log.ForceInfof("Running grpc command '%s'", action.Name)

	res, err := action.Execute()
	if err != nil {
		return nil, err
	}

	return &pb.RunActionResponse{
		Action:  res.Action.Name,
		Command: res.Action.Command,
		Stdout:  res.Stdout,
		Code:    int64(res.Code),
	}, nil
}
