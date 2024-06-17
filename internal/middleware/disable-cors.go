package middleware

import "net/http"

func DisableCORS(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(`Access-Control-Allow-Origin`, `*`)
		w.Header().Set(`Access-Control-Allow-Methods`, `GET, POST, PATCH, PUT, DELETE, OPTIONS`)
		w.Header().Set(`Access-Control-Allow-Headers`, `*`)
		h.ServeHTTP(w, r)
	})
}
