package ext

import (
	"github.com/ffip/tata/api/user/cfg"
	"github.com/ffip/tata/api/user/route"
	http "github.com/ffip/tata/lib/net/http/server/service"
	"github.com/sirupsen/logrus"
)

// ListenAndServe ==> Tita ListenAndServe.
func ListenAndServe(log *logrus.Logger) {
	ws := http.Server{
		Addr:  cfg.Web.HttpAddr,
		Crt:   cfg.Web.HttpsCrt,
		Key:   cfg.Web.HttpsKey,
		Route: route.Path(),
	}

	ws.ListenAndServe(log)
}
