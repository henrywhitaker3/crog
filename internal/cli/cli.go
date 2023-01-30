package cli

import (
	"os"

	"github.com/pterm/pterm"
)

func SingleChoice(options []string) (string, error) {
	return pterm.DefaultInteractiveSelect.WithOptions(options).Show()
}

func ErrorExit(err error) {
	pterm.Error.Printfln("%s", err)
	os.Exit(1)
}

func Printfln(format string, a ...any) {
	pterm.Info.Printfln(format, a...)
}
