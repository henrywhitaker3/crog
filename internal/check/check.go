package check

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/google/shlex"
)

type Check struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Id      string `yaml:"id"`
	Code    int    `yaml:"code"`
	Cron    string `yaml:"cron"`
}

func (c *Check) Execute() error {
	c.startCheck()

	code, _ := c.runCommand()

	if code != c.Code {
		c.failCheck()
		return fmt.Errorf("check failed, expected status %d, got %d", c.Code, code)
	}

	c.completeCheck()

	return nil
}

func (c *Check) runCommand() (int, string) {
	args, _ := shlex.Split(c.Command)
	bin := args[0]
	args = args[1:]

	cmd := exec.Command(bin, args...)
	out, err := cmd.Output()

	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode(), string(out)
	}

	return 0, string(out)
}

func (c *Check) startCheck() error {
	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/start", c.Id))

	return err
}
func (c *Check) completeCheck() error {
	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s", c.Id))

	return err
}
func (c *Check) failCheck() error {
	_, err := http.Get(fmt.Sprintf("https://hc-ping.com/%s/fail", c.Id))

	return err
}
