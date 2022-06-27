package handlers

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func initLogConfiguration() {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02T15:04:05+0000"
	customFormatter.FullTimestamp = true

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	log.SetFormatter(customFormatter)
	log.SetLevel(logLevel)

	// The log level is the hierarchy.
	log.Trace("Successfully initial log Trace")
	log.Debug("Successfully initial log Debug")
	log.Info("Successfully initial log Info")
	log.Warn("Successfully initial log Warn")
	log.Error("Successfully initial log Error")
	// log.Fatal("Bye.")         // Calls os.Exit(1) after logging
	// log.Panic("I'm bailing.") // Calls panic() after logging
	Log = log
}

func Trace() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s.%d", f.Name(), line)
}

func HandleErr(err error, message string) {
	if err != nil {
		Log.Fatalf("%s: %v", message, err)
	}
}
