package action

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/google/shlex"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/validation"
)

type Action struct {
	Name    string `yaml:"name" required:"true"`
	Command string `yaml:"command" required:"true"`
	Code    int    `yaml:"code" default:"0"`
	Cron    string `yaml:"cron" default:"* * * * *"`
	On      On     `yaml:"on" requred:"true"`
}

func (a *Action) Execute() error {
	a.LogInfo("Executing check")
	a.start()

	code, _ := a.runCommand()

	if code != a.Code {
		a.LogError(fmt.Sprintf("Check failed - expected status %d, got %d", a.Code, code))
		a.fail()
		return fmt.Errorf("check failed - expected status %d, got %d", a.Code, code)
	}

	a.LogInfo("Check passed")

	a.success()

	return nil
}

func (a *Action) LogInfo(value string) {
	log.Infof("[%s] %s", a.Name, value)
}

func (a *Action) LogError(value string) {
	log.Errorf("[%s] %s", a.Name, value)
}

func (a *Action) runCommand() (int, string) {
	a.LogInfo(fmt.Sprintf("Running command: '%s'", a.Command))
	args, _ := shlex.Split(a.Command)
	bin := args[0]
	args = args[1:]
	a.LogInfo(fmt.Sprintf("executable: %s", bin))
	a.LogInfo(fmt.Sprintf("args: %v", args))

	cmd := exec.Command(bin, args...)
	out, err := cmd.Output()

	exitCode := 0

	if exitError, ok := err.(*exec.ExitError); ok {
		exitCode = exitError.ExitCode()
	}

	a.LogInfo(fmt.Sprintf("Got exit code %d", exitCode))
	a.LogInfo(fmt.Sprintf("Got stdout:\n%s", string(out)))

	return exitCode, string(out)
}

func (a *Action) start() error {
	if a.On.Start == "" {
		return nil
	}

	a.LogInfo("Sending start request for check")
	_, err := http.Get(a.On.Start)

	return err
}
func (a *Action) success() error {
	a.LogInfo("Sending success request for check")

	_, err := http.Get(a.On.Success)

	return err
}
func (a *Action) fail() error {
	if a.On.Failure == "" {
		return nil
	}

	a.LogInfo("Sending failure request for check")
	_, err := http.Get(a.On.Failure)

	return err
}

func (a *Action) Validate() error {
	if err := a.On.Validate(); err != nil {
		return err
	}
	return validation.Validate(a)
}
