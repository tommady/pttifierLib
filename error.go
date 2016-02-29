package pttifierLib

import (
	"fmt"
	"time"
)

type PttiferErr struct {
	When time.Time
	What string
	Err  error
}

func (e PttiferErr) Error() string {
	return fmt.Sprintf("%v -> %v -> %v", e.When, e.What, e.Err)
}

func ReportError(what string, err error) error {
	return PttiferErr{
		time.Now(),
		what,
		err,
	}
}
