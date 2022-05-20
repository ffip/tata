package ext

import (
	"bitbucket.org/pwq/tata/api/user/cfg"
	"bitbucket.org/pwq/tata/api/user/route"
	http "bitbucket.org/pwq/tata/lib/net/http/server/service"
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
