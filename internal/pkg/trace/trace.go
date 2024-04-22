// Package trace -.
package trace

import (
	"fmt"
	"runtime"
)

func funcName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "?"
	}

	return fn.Name()
}

// Wrap -.
func Wrap(err error) error {
	return fmt.Errorf(funcName(2)+": %w", err) //nolint:gomnd
}
