package cli

import (
	"github.com/pterm/pterm"
)

func SingleChoice(options []string) (string, error) {
	return pterm.DefaultInteractiveSelect.WithOptions(options).Show()
}

func ErrorExit(err error) {
	pterm.Error.Printfln("%s", err)
}

func Printfln(format string, a ...any) {
	pterm.Info.Printfln(format, a...)
}
