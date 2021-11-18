package end_errors

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	allErrors = map[string]Error{}
)

type EndError struct {
	AllErrors map[string]Error
}

type Error struct {
	Message      string
	FunctionName string
}

func (e Error) Error() string {
	return fmt.Sprintf("* %s: %s\n", e.FunctionName, e.Message)
}

func GetInstance() *EndError {
	return &EndError{
		AllErrors: allErrors,
	}
}

func (ee *EndError) Count() int {
	return len(ee.AllErrors)
}

func (ee *EndError) PrintErrors() {
	if len(ee.AllErrors) == 0 {
		return
	}

	fmt.Printf("\n\n%d error(s) were found while running command:\n", len(ee.AllErrors))
	for _, singleError := range ee.AllErrors {
		fmt.Print(singleError.Error())
		fmt.Println("")
	}
}

func (ee *EndError) AddErrorf(format string, a ...interface{}) {
	var stackPrefix string
	if _, file, line, ok := runtime.Caller(1); ok {
		stackPrefix = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	newError := Error{
		Message:      fmt.Sprintf(format, a...),
		FunctionName: stackPrefix,
	}

	_, ok := ee.AllErrors[newError.Message]
	if !ok {
		ee.AllErrors[newError.Message] = newError
	}
}
