package logging

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"runtime"
)

func init() {
	file, err := os.OpenFile("transaction.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed to open log file %v", err)
	}

	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	formatter := &logrus.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", ""
		},
	}
	logrus.SetFormatter(formatter)
	logrus.SetOutput(file)
}
