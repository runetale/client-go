package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupkeyHandlerManager interface {
	SetupKey(w http.ResponseWriter, r *http.Request)
}

type SetupkeyHandler struct {
	redis        *redis.RedisClient
	config       *config.Config
	accountStore *redis.AccountStore
	serverStore  *store.ServerStore
}

func NewSetupKeyHanlder(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore,
) *SetupkeyHandler {
	return &SetupkeyHandler{
		redis:        r,
		config:       config,
		accountStore: account,
		serverStore:  server,
	}
}

type SetupKeyRequest struct {
	Group      string
	Job        string
	Permission string
}

func (h *SetupkeyHandler) SetupKey(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req SetupKeyRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sub := r.Header.Get("sub")
		setupKey, err := key.NewSetupKey(sub, req.Group, req.Job, key.PermissionType(req.Permission))
		if err != nil {
			http.Error(w, "failed to create setupkey", http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(setupKey)
	case http.MethodDelete:
		log.Println("delete setupkey")
	default:
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	}
}
