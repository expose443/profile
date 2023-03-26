package transport

import "net/http"

func ErrorHandler(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
	w.WriteHeader(status)
}
