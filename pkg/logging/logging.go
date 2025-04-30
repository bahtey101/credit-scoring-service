package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

func SetLogging(logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(UTCFormatter{&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			wd, err := os.Getwd()
			file := f.File
			if err == nil {
				if rel, err := filepath.Rel(wd, f.File); err == nil {
					file = rel
				}
			}
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", file, f.Line)
		},
	}})

	logrus.Info("logging setting set")
}
