package log

import (
	"github.com/sirupsen/logrus"
)

// NewLogger ==> Creates a new logrus instance with the provided level
func NewLogger(level logrus.Level) (log *logrus.Logger) {
	log = logrus.New()
	// log.SetFormatter(new(prefixed.TextFormatter))
	log.SetLevel(level)

	return
}
