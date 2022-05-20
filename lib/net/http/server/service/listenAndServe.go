package service

import (
	_ "expvar"
	"fmt"
	"net"
	"net/http"

	using "bitbucket.org/pwq/tata/lib/net/utils"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

// Server 		==> HTTP服务端
type Server struct {
	Route *chi.Mux
	Addr  string
	Crt   string
	Key   string
}

// New server
func New() (s *Server) {
	s = &Server{}
	return
}

// ListenAndServe 		==> 启动http服务
func (h *Server) ListenAndServe(log *logrus.Logger) (err error) {
	log.Info("API-CHI: ListenAndServe - Listening ...")

	httpAddr, err := net.ResolveTCPAddr("tcp", h.Addr)
	if err != nil {
		err = fmt.Errorf("%s", "The listening address is incorrect or empty")
		addr, _ := using.GetRandomTCPAddress(httpAddr.Port, 65535)
		httpAddr.Port = addr.Port
		return
	}

	switch {
	case h.Crt != "" && h.Key != "":
		log.Info(fmt.Sprintf("API-CHI: using %s%s\n", "https://", httpAddr.String()))
		err = http.ListenAndServeTLS(httpAddr.String(), h.Crt, h.Key, h.Route)
	default:
		log.Info(fmt.Sprintf("API-CHI: using %s%s\n", "http://", httpAddr.String()))
		err = http.ListenAndServe(httpAddr.String(), h.Route)
	}

	if err != nil {
		log.Info(fmt.Sprintf("API-CHI: %s\n", err))
		return
	}
	return
}
