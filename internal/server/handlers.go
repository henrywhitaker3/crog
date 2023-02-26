package server

import (
	"context"

	"github.com/henrywhitaker3/crog/internal/action"
	"github.com/henrywhitaker3/crog/internal/event"
	"github.com/henrywhitaker3/crog/internal/pb"
	"github.com/henrywhitaker3/events"
)

func (s Server) List(context.Context, *pb.ListActionsRequest) (*pb.ListActionsResponse, error) {
	acts := []*pb.Action{}

	for _, a := range s.cfg.Actions {
		acts = append(acts, actionToPbAction(a))
	}

	return &pb.ListActionsResponse{Actions: acts}, nil
}

func (s Server) Run(ctx context.Context, req *pb.RunActionRequest) (*pb.RunActionResponse, error) {
	action, err := s.cfg.GetAction(req.Action)
	if err != nil {
		return nil, err
	}

	event.Emit(events.Event{Tag: "ActionPreflight", Data: action})
	res := action.Execute()
	event.Emit(events.Event{Tag: "Result", Data: res})

	errS := ""
	if res.GetErr() != nil {
		errS = res.GetErr().Error()
	}

	return &pb.RunActionResponse{
		Action: actionToPbAction(action),
		Err:    errS,
		Stdout: res.GetStdout(),
		Code:   int64(res.GetCode()),
	}, nil
}

func actionToPbAction(action *action.Action) *pb.Action {
	return &pb.Action{
		Name:    action.Name,
		Command: action.Command,
		Cron:    action.Cron,
		Code:    int64(action.Code),
		When: &pb.ActionWhen{
			Start:   action.On.Start,
			Success: action.On.Success,
			Failure: action.On.Failure,
		},
	}
}
