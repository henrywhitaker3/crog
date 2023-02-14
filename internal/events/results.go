package events

import (
	"fmt"

	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/log"
)

var ResultHandler eventHandler

type Result struct {
	Result domain.Result
}

type ResultLogger struct{}

func (a *ResultLogger) Handle(e domain.Event) error {
	ev := e.(Result)
	fmt.Println("here")
	log.LogResult(ev.Result)
	return nil
}
