package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Notch-Technologies/wizy/client"
)

func adminMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	sub := r.Header.Get("sub")

	if !strings.HasPrefix(sub, "auth0") {
    	http.Error(w, "invalid value", http.StatusBadRequest)
	}

	accessToken, err := client.GetAuth0ManagementAccessToken()
	if err != nil {
    	http.Error(w, fmt.Sprintf("%s", err.Error()), http.StatusBadRequest)
	}

	isAdmin, err := client.IsAdmin(sub, accessToken)
	if err != nil || !isAdmin {
    	http.Error(w, fmt.Sprintf("%s", err.Error()), http.StatusBadRequest)
	}

    next.ServeHTTP(w, r)
  })
}
