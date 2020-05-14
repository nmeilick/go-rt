package rt

import (
	"fmt"
	"runtime"
)

// CallerInfo holds info about a caller.
type CallerInfo struct {
	Name string
	File string
	Line int
}

func (ci *CallerInfo) String() string {
	if ci == nil {
		return "<unknown caller>"
	}
	return fmt.Sprintf("%s (%s:%d)", ci.Name, ci.File, ci.Line)
}

// NewCallerInfo creates a new CallerInfo initialized with the given values.
func NewCallerInfo(name, file string, line int) *CallerInfo {
	return &CallerInfo{
		Name: name,
		File: file,
		Line: line,
	}
}

// GetCaller returns info about a calling function, with skip indicating
// how many stack frames to skip (0 = immediate caller, 1 = caller's caller, etc.)
// If the info cannot be retrieved, nil is returned.
func GetCaller(skip int) *CallerInfo {
	pc := make([]uintptr, 1)
	if runtime.Callers(3+skip, pc) > 0 {
		if f := runtime.FuncForPC(pc[0] - 1); f != nil {
			file, line := f.FileLine(pc[0] - 1)
			return NewCallerInfo(f.Name(), file, line)
		}
	}
	return nil
}

// Caller returns caller info about the immediate parent caller of the
// calling function.
func Caller() *CallerInfo {
	return GetCaller(0)
}
