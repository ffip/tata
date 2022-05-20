package service

import (
	"testing"

	"bitbucket.org/pwq/tata/lib/log"
	"github.com/sirupsen/logrus"
)

func TestEcho(t *testing.T) {
	logLevel := logrus.InfoLevel
	log := log.NewLogger(logLevel)
	var h Server
	h.Addr = "127.0.0.1:1010"
	h.ListenAndServe(log)
}
