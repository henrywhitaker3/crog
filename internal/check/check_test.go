package check_test

import (
	"testing"

	"github.com/henrywhitaker3/crog/internal/check"
)

func TestItReturnsAnErrorIfTheNameIsNotSet(t *testing.T) {
	check := check.Check{
		Id:      "apple",
		Command: "bongo",
	}

	if err := check.Validate(); err == nil {
		t.Fail()
	}
}

func TestItReturnsAnErrorIfTheCommandIsNotSet(t *testing.T) {
	check := check.Check{
		Id:   "apple",
		Name: "bongo",
	}

	if err := check.Validate(); err == nil {
		t.Fail()
	}
}

func TestItReturnsAnErrorIfTheIdIsNotSet(t *testing.T) {
	check := check.Check{
		Command: "apple",
		Name:    "bongo",
	}

	if err := check.Validate(); err == nil {
		t.Fail()
	}
}

func TestItSetsTheDefaultValueForCron(t *testing.T) {
	check := check.Check{
		Command: "apple",
		Name:    "bongo",
		Id:      "bingo",
	}

	if err := check.Validate(); err != nil {
		t.Fail()
	}

	if check.Cron != "* * * * *" {
		t.Fail()
	}
}

func TestItSetsTheDefaultValueForCode(t *testing.T) {
	check := check.Check{
		Command: "apple",
		Name:    "bongo",
		Id:      "bingo",
	}

	if err := check.Validate(); err != nil {
		t.Fail()
	}

	if check.Code != 0 {
		t.Fail()
	}
}
