package route

import (
	"net/http"

	"bitbucket.org/pwq/tata/api/user/model"
	"github.com/go-chi/chi"
)

// Path ==> Route Path
func Path() (r *chi.Mux) {
	r = chi.NewRouter()

	r.Post("/new", model.New)
	r.Post("/info", model.Info)
	r.Post("/find", model.Find)
	r.Post("/login", model.Login)
	r.Post("/update", model.Update)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json; charset=utf-8")
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		_, _ = writer.Write([]byte("{'info':'Cato - Tita - Console'}"))
	})
	return
}
