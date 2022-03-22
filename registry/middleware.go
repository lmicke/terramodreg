package registry

import (
	"log"
	"net/http"
)

func LoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] requested uri %s from %s", r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
