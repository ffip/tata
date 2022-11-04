package ext

import (
	"github.com/ffip/tata/api/sms/cfg"
	"github.com/ffip/tata/api/sms/route"
	log "github.com/ffip/tata/lib/log"
	http "github.com/ffip/tata/lib/net/http/server/service"
)

// ListenAndServe ==> Tita ListenAndServe.
func ListenAndServe(log *log.Logger) {
	ws := http.Server{
		Addr:  cfg.Web.HttpAddr,
		Crt:   cfg.Web.HttpsCrt,
		Key:   cfg.Web.HttpsKey,
		Route: route.Path(),
	}

	ws.ListenAndServe(log)
}
