package action

import (
	"context"
	"io/ioutil"
	"net/http"
	"os/exec"

	"github.com/henrywhitaker3/crog/internal/circuits"
	"github.com/henrywhitaker3/crog/internal/domain"
	"github.com/henrywhitaker3/crog/internal/validation"
)

type Action struct {
	Name    string `yaml:"name" required:"true"`
	Command string `yaml:"command" required:"true"`
	Code    int    `yaml:"code" default:"0"`
	Cron    string `yaml:"cron" default:"* * * * *"`
	Tries   int    `yaml:"tries" default:"1"`
	On      On     `yaml:"when"`
}

var client *http.Client

func init() {
	client = &http.Client{}
}

func (a *Action) Execute() domain.Result {
	// TODO: add results struct for start/failure/success actions
	a.start()

	tries := 0
	retry := circuits.Retry(func(ctx context.Context) (any, error) {
		tries++
		code, out := a.runCommand()

		res := Result{
			Action: a,
			Code:   code,
			Stdout: out,
		}

		if code != a.Code {
			a.fail()
			res.Err = ActionFailed{Expected: a.Code, Actual: code}
			return res, res.Err
		}

		a.success()
		return res, nil
	}, a.Tries)

	r, _ := retry(context.Background())
	res := r.(Result)
	res.Tries = tries

	return res
}

func (a *Action) runCommand() (int, string) {
	cmd := exec.Command("/bin/bash", "-c", a.Command)
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

func (a Action) GetName() string {
	return a.Name
}

func (a Action) GetCommand() string {
	return a.Command
}

func (a Action) GetCron() string {
	return a.Cron
}

func (a Action) GetTries() int {
	return a.Tries
}
