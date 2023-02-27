package validation

import "testing"

func TestItErrorsWhenRequiredBoolIsNotSet(t *testing.T) {
	s := &struct {
		Bongo bool `required:"true" default:"true"`
	}{}
	if err := Validate(s); err == nil {
		t.Fail()
	}
}

func TestItDoesntErrorsWhenRequiredBoolIsSet(t *testing.T) {
	s := &struct {
		Bongo bool `required:"true" default:"true"`
	}{Bongo: true}
	if err := Validate(s); err != nil {
		t.Fail()
	}
}

func TestItSetsADefaultBoolValue(t *testing.T) {
	s := &struct {
		Bongo bool `default:"true"`
	}{}
	Validate(s)

	if s.Bongo != true {
		t.Fail()
	}
}

func TestItErrorsWhenRequiredIntIsNotSet(t *testing.T) {
	s := &struct {
		Bongo int `required:"true" default:"1"`
	}{}
	if err := Validate(s); err == nil {
		t.Fail()
	}
}

func TestItDoesntErrorsWhenRequiredIntIsSet(t *testing.T) {
	s := &struct {
		Bongo int `required:"true" default:"1"`
	}{Bongo: 2}
	if err := Validate(s); err != nil {
		t.Fail()
	}
}

func TestItSetsADefaultIntValue(t *testing.T) {
	s := &struct {
		Bongo int `default:"1"`
	}{}
	Validate(s)

	if s.Bongo != 1 {
		t.Fail()
	}
}

func TestItErrorsWhenRequiredStringIsNotSet(t *testing.T) {
	s := &struct {
		Bongo string `required:"true" default:"pie"`
	}{}
	if err := Validate(s); err == nil {
		t.Fail()
	}
}

func TestItDoesntErrorsWhenRequiredStringIsSet(t *testing.T) {
	s := &struct {
		Bongo string `required:"true" default:"pie"`
	}{Bongo: "fish"}
	if err := Validate(s); err != nil {
		t.Fail()
	}
}

func TestItSetsADefaultStringValue(t *testing.T) {
	s := &struct {
		Bongo string `default:"pie"`
	}{}
	Validate(s)

	if s.Bongo != "pie" {
		t.Fail()
	}
}
