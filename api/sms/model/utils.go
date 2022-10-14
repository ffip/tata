package model

import (
	"net/http"
	"strconv"

	"github.com/ffip/tata/lib/text"
)

func getValueInt(writer http.ResponseWriter, request *http.Request, item string) (out int, err error) {
	out, err = strconv.Atoi(request.FormValue(item))
	if err != nil {
		fail(writer, err)
		return
	}
	return
}

func fail(writer http.ResponseWriter, err error) {
	writer.WriteHeader(400)
	_, _ = writer.Write(text.Atob(text.Mgr("{\"err\":\"", err.Error(), "\"}")))
}

func ok(writer http.ResponseWriter) {
	_, _ = writer.Write(text.Atob(text.Mgr("{\"status\":true}")))
}

func ref(writer http.ResponseWriter, err error, data string) {
	if err != nil {
		fail(writer, err)
		return
	}
	_, _ = writer.Write(text.Atob(text.Mgr("{\"data\":\"", data, "\"}")))
}
