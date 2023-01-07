package action_test

import (
	"testing"

	"github.com/henrywhitaker3/crog/internal/action"
)

func TestItErrorsWhenSuccessIsNotSet(t *testing.T) {
	on := action.On{}

	if err := on.Validate(); err == nil {
		t.Fail()
	}
}

func TestItDoesntErrorWhenSuccessIsSet(t *testing.T) {
	on := action.On{Success: "bongo"}

	if err := on.Validate(); err != nil {
		t.Fail()
	}
}
