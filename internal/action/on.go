package action

import "github.com/henrywhitaker3/crog/internal/validation"

type On struct {
	Start   string `yaml:"start"`
	Success string `yaml:"success" required:"true"`
	Failure string `yaml:"failure"`
}

func (o *On) Validate() error {
	return validation.Validate(o)
}
