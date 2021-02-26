package util

import (
	"os"
	"runtime"

	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

//Initialization for logs
func init() {
	logger = log.New()
	logger.SetLevel(log.TraceLevel)
	logger.Formatter = &log.TextFormatter{}
	log.SetOutput(os.Stdout)
}

//Logger with fields
func Logger() *log.Entry {
	var depth = 1
	function, file, line, _ := runtime.Caller(depth)
	functionObject := runtime.FuncForPC(function)
	entry := logger.WithFields(log.Fields{
		"file":     file,
		"function": functionObject.Name(),
		"line":     line,
	})

	logger.SetOutput(os.Stdout)
	return entry

}
