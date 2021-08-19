package utils

import (
	"net/http"
)

var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		Respond(w, Message(http.StatusNotFound, "Resource was not found"))
		next.ServeHTTP(w, r)
	})
}
