// Package trace -.
package trace

import (
	"fmt"
	"runtime"
	"strings"
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

	funcNameWithModule := fn.Name()
	funcNameWithModuleSplit := strings.Split(funcNameWithModule, "/")
	funcName := funcNameWithModuleSplit[len(funcNameWithModuleSplit)-1]

	return funcName
}

// Wrap -.
func Wrap(err error) error {
	return fmt.Errorf(funcName(2)+": %w", err) //nolint:gomnd
}
