package log

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var logger = logrus.New()

func PrintWaLog(c echo.Context) waLog.Logger {
	return GetLogger(c)
}

func PrintOurLogger(c echo.Context) *OurLogger {
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
		DisableColors:   false,
		ForceColors:     true,
	}
	return GetLogger(c)
}

func Print(c echo.Context) *logrus.Entry {
	return PrintOurLogger(c).Entry
}

type OurLogger struct {
	*logrus.Entry
}

func (l OurLogger) Sub(module string) waLog.Logger {
	return OurLogger{Entry: l.Entry.WithField("module", module)}
}

func NewLogger() *OurLogger {
	log := logger.WithFields(logrus.Fields{})
	return &OurLogger{Entry: log}
}

func GetLogger(c echo.Context) *OurLogger {
	if c == nil {
		return &OurLogger{Entry: logger.WithFields(logrus.Fields{})}
	}
	return &OurLogger{
		Entry: logger.WithFields(logrus.Fields{
			"remote_ip": c.Request().RemoteAddr,
			"method":    c.Request().Method,
			"uri":       c.Request().URL.String(),
		}),
	}
}
