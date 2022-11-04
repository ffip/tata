package service

import (
	"testing"

	log "github.com/ffip/tata/lib/log"
)

func TestEcho(t *testing.T) {
	loger := log.Logger{}
	var h Server
	h.Addr = "127.0.0.1:1010"
	h.ListenAndServe(&loger)
}
