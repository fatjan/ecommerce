package logger

import (
	"runtime"
	"strings"
)

// LogFormat this is for log format
func LogFormat(data interface{}, msg string, severity string) {
	fileName, funcName := Trace()

	Logger.Info().
		Str("module", fileName).
		Str("operation", funcName).
		Str("message", msg).
		Str("severity", severity).
		Msgf("data : %s", data)
}

// Trace func to get current file name and function name
func Trace() (string, string) {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, _ := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())

	var fileName, funcName string
	filePath := strings.Split(file, "magellan")
	if len(filePath) > 1 {
		fileName = filePath[1][1:]
	}

	funcPath := strings.Split(f.Name(), ".")
	funcName = funcPath[len(funcPath)-1]
	return fileName, funcName
}
