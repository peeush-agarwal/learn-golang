package middlewares

import "net/http"

func ContentTypeJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		h.ServeHTTP(w, r)
	})
}
