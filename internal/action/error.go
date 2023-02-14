package action

import "fmt"

type ActionFailed struct {
	Expected int
	Actual   int
}

func (a ActionFailed) Error() string {
	return fmt.Sprintf("expected code %d, got %d", a.Expected, a.Actual)
}
