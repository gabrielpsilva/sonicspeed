package restserver

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func GlobalLogrusFormatJSON(loglevel logrus.Level) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(loglevel)
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "06-01-02 15:04:05.000",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			_, funcname := path.Split(f.Function)
			return funcname, fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
}

func GlobalLogrusFormatTEXT(loglevel logrus.Level) {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(loglevel)
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:    false,
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		ForceColors:      true,
		TimestampFormat:  "06-01-02 15:04:05.0000",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {

			_, filename := path.Split(f.File)

			return fmt.Sprintf("%s:%-50d", filename, f.Line)[:30] + "|", " "
		},
	})
}
