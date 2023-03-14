package event

import (
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/events"
)

type ServerStartedLogger struct{}

func (s *ServerStartedLogger) Handle(e events.Event) error {
	log.Log.Infof("Starting grpc server on %s", e.Data)
	return nil
}

type GrpcRequestLogger struct{}

func (g *GrpcRequestLogger) Handle(e events.Event) error {
	// TODO: implement this
	return nil
}
