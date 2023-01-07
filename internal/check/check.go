package check

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/google/shlex"
	"github.com/henrywhitaker3/crog/internal/log"
	"github.com/henrywhitaker3/crog/internal/validation"
)

type Check struct {
	Name    string `yaml:"name" required:"true"`
	Command string `yaml:"command" required:"true"`
	Id      string `yaml:"id" required:"true"`
	Code    int    `yaml:"code" default:"0"`
	Cron    string `yaml:"cron" default:"* * * * *"`
}

func (c *Check) Execute() error {
	c.LogInfo("Executing check")
	c.startCheck()

	code, _ := c.runCommand()

	if code != c.Code {
		c.LogError(fmt.Sprintf("Check failed - expected status %d, got %d", c.Code, code))
		c.failCheck()
		return fmt.Errorf("check failed - expected status %d, got %d", c.Code, code)
	}

	c.LogInfo("Check passed")

	c.completeCheck()

	return nil
}

func (c *Check) LogInfo(value string) {
	log.Infof("[%s] %s", c.Name, value)
}

func (c *Check) LogError(value string) {
	log.Errorf("[%s] %s", c.Name, value)
}

func (c *Check) runCommand() (int, string) {
	c.LogInfo(fmt.Sprintf("Running command: '%s'", c.Command))
	args, _ := shlex.Split(c.Command)
	bin := args[0]
	args = args[1:]
	c.LogInfo(fmt.Sprintf("executable: %s", bin))
	c.LogInfo(fmt.Sprintf("args: %v", args))

	cmd := exec.Command(bin, args...)
	out, err := cmd.Output()

	exitCode := 0

	if exitError, ok := err.(*exec.ExitError); ok {
		exitCode = exitError.ExitCode()
	}

	c.LogInfo(fmt.Sprintf("Got exit code %d", exitCode))
	c.LogInfo(fmt.Sprintf("Got stdout:\n%s", string(out)))

	return exitCode, string(out)
}

func (c *Check) startCheck() error {
	c.LogInfo("Sending start GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/start", c.Id))

	return err
}
func (c *Check) completeCheck() error {
	c.LogInfo("Sending finish GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s", c.Id))

	return err
}
func (c *Check) failCheck() error {
	c.LogInfo("Sending fail GET request for check")

	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/fail", c.Id))

	return err
}

func (c *Check) Validate() error {
	return validation.Validate(c)
}
