package config

import (
	"testing"

	"github.com/henrywhitaker3/crog/internal/action"
)

func TestItGetsAnActionByName(t *testing.T) {
	cfg := Config{
		Actions: []*action.Action{
			{Name: "bongo"},
		},
	}

	if _, err := cfg.GetAction("bongo"); err != nil {
		t.Error(err)
	}
}

func TestItErrorsWhenThereIsNoAction(t *testing.T) {
	cfg := Config{
		Actions: []*action.Action{},
	}

	if _, err := cfg.GetAction("bongo"); err == nil {
		t.Error("there is no error")
	}
}

func TestItGetsARemoteByName(t *testing.T) {
	cfg := Config{
		Remotes: []Remote{
			{Name: "bongo"},
		},
	}

	if _, err := cfg.GetRemote("bongo"); err != nil {
		t.Error(err)
	}
}

func TestItErrorsWhenThereIsNoRemote(t *testing.T) {
	cfg := Config{
		Remotes: []Remote{},
	}

	if _, err := cfg.GetRemote("bongo"); err == nil {
		t.Error("there is no error")
	}
}
