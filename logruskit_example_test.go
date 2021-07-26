package logruskit_test

import (
	"os"

	"github.com/colega/logruskit"

	"github.com/go-kit/log"
)

func ExampleLogruskit_Simple() {
	logkitLogger := log.NewLogfmtLogger(os.Stdout)
	logrusLogger := logruskit.NewLogger(logkitLogger)

	logrusLogger.WithField("hello", "world").Infof("From: %s", "logrus")
	// Output:
	// level=info hello=world msg="From: logrus"
}
