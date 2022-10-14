package route

import (
	"net/http"

	"github.com/ffip/tata/api/sms/model"
	"github.com/go-chi/chi"
)

// Path ==> Route Path
func Path() (r *chi.Mux) {
	r = chi.NewRouter()

	r.Post("/new", model.New)
	r.Post("/used", model.Used)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		_, _ = writer.Write([]byte("{'info':'Cato - Tita - Console'}"))
	})
	return
}
