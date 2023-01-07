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
	Id      string `yaml:"id" required:"true"`
	Code    int    `yaml:"code" default:"0"`
	Cron    string `yaml:"cron" default:"* * * * *"`
}

func (a *Action) Execute() error {
	a.LogInfo("Executing check")
	a.startCheck()

	code, _ := a.runCommand()

	if code != a.Code {
		a.LogError(fmt.Sprintf("Check failed - expected status %d, got %d", a.Code, code))
		a.failCheck()
		return fmt.Errorf("check failed - expected status %d, got %d", a.Code, code)
	}

	a.LogInfo("Check passed")

	a.completeCheck()

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

func (a *Action) startCheck() error {
	a.LogInfo("Sending start GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/start", a.Id))

	return err
}
func (a *Action) completeCheck() error {
	a.LogInfo("Sending finish GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s", a.Id))

	return err
}
func (a *Action) failCheck() error {
	a.LogInfo("Sending fail GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/fail", a.Id))

	return err
}

func (a *Action) Validate() error {
	return validation.Validate(a)
}
