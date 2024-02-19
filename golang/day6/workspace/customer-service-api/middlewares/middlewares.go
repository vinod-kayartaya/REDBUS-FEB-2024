package middlewares

import (
	"log"
	"net/http"
)

func RejectNonJsonRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Accept"][0] != "application/json" {
			w.WriteHeader(http.StatusNotAcceptable)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s (%s)", r.Method, r.URL, r.RemoteAddr)

		next.ServeHTTP(w, r)
	})
}
