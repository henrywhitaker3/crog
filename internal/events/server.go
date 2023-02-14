package events

import (
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

var ServerStartedHandler eventHandler

type ServerStarted struct {
	Address string
}

type ServerStartedLogger struct{}

func (s *ServerStartedLogger) Handle(e domain.Event) error {
	ev := e.(ServerStarted)
	log.ForceInfof("Starting grpc server on %s", ev.Address)
	return nil
}
