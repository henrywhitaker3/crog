package events

import (
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

type ServerStarted struct {
	Address string
}

type ServerStartedLogger struct{}

func (s *ServerStartedLogger) Handle(e domain.Event) error {
	ev := e.(ServerStarted)
	log.Log.ForceInfof("Starting grpc server on %s", ev.Address)
	return nil
}

type GrpcRequest struct{}

type GrpcRequestLogger struct{}

func (g *GrpcRequestLogger) Handle(e domain.Event) error {
	// TODO: implement this
	return nil
}
