package api

import (
	"log"
	"net/http"
)

func setupKey(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			log.Println("post setupkey")

		case http.MethodDelete:
			log.Println("delete Setupkey")
		default:
    		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	}
}
