package action_test

import (
	"testing"

	"github.com/henrywhitaker3/crog/internal/action"
)

func TestItReturnsAnErrorIfTheNameIsNotSet(t *testing.T) {
	check := action.Action{
		Command: "bongo",
	}

	if err := check.Validate(); err == nil {
		t.Fail()
	}
}

func TestItReturnsAnErrorIfTheCommandIsNotSet(t *testing.T) {
	check := action.Action{
		Name: "bongo",
	}

	if err := check.Validate(); err == nil {
		t.Fail()
	}
}

func TestItSetsTheDefaultValueForCron(t *testing.T) {
	check := action.Action{
		Command: "apple",
		Name:    "bongo",
	}

	if err := check.Validate(); err != nil {
		t.Fail()
	}

	if check.Cron != "* * * * *" {
		t.Fail()
	}
}

func TestItSetsTheDefaultValueForCode(t *testing.T) {
	check := action.Action{
		Command: "apple",
		Name:    "bongo",
	}

	if err := check.Validate(); err != nil {
		t.Fail()
	}

	if check.Code != 0 {
		t.Fail()
	}
}
