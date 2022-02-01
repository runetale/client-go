package api

import (
	"net/http"
	"strings"
)

func adminMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sub := r.Header.Get("sub")

	if !strings.HasPrefix(sub, "auth0") {
    	http.Error(w, "invalid value", http.StatusBadRequest)
	}

    next.ServeHTTP(w, r)
  })
}
