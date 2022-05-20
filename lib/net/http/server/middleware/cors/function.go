package cors

import "net/http"

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	_, _ = w.Write([]byte("{}"))
}
