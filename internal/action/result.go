package action

import "github.com/henrywhitaker3/crog/internal/domain"

type Result struct {
	Action *Action
	Err    error
	Code   int
	Stdout string
}

func (r Result) GetErr() error {
	return r.Err
}

func (r Result) GetCode() int {
	return r.Code
}

func (r Result) GetStdout() string {
	return r.Stdout
}

func (r Result) GetAction() domain.Action {
	return r.Action
}
