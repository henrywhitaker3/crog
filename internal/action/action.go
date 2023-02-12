package action

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/google/shlex"
	"github.com/henrywhitaker3/crog/internal/validation"
)

type Action struct {
	Name    string `yaml:"name" required:"true"`
	Command string `yaml:"command" required:"true"`
	Code    int    `yaml:"code" default:"0"`
	Cron    string `yaml:"cron" default:"* * * * *"`
	On      On     `yaml:"when"`
}

type Result struct {
	Action *Action
	Err    error
	Code   int
	Stdout string
}

var client *http.Client

func init() {
	client = &http.Client{}
}

func (a *Action) Execute() *Result {
	// TODO: add results struct for start/failure/success actions
	a.start()

	code, out := a.runCommand()

	res := &Result{
		Action: a,
		Code:   code,
		Stdout: out,
	}

	if code != a.Code {
		a.fail()
		res.Err = fmt.Errorf("check failed - expected status %d, got %d", a.Code, code)
		return res
	}

	a.success()

	return res
}

func (a *Action) runCommand() (int, string) {
	args, _ := shlex.Split(a.Command)
	bin := args[0]
	args = args[1:]
	cmd := exec.Command(bin, args...)
	out, err := cmd.Output()

	exitCode := 0
	if exitError, ok := err.(*exec.ExitError); ok {
		exitCode = exitError.ExitCode()
	}

	return exitCode, string(out)
}

func (a *Action) start() error {
	if a.On.Start == "" {
		return nil
	}

	return a.request(a.On.Start)
}

func (a *Action) success() error {
	if a.On.Success == "" {
		return nil
	}

	return a.request(a.On.Success)
}

func (a *Action) fail() error {
	if a.On.Failure == "" {
		return nil
	}

	return a.request(a.On.Failure)
}

func (a *Action) request(url string) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Crog")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (a *Action) Validate() error {
	if err := a.On.Validate(); err != nil {
		return err
	}
	return validation.Validate(a)
}
