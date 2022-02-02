package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/model"
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
	user        *redis.UserStore
}

func NewSetupKeyHanlder(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, user *redis.UserStore,
) *SetupkeyHandler {
	return &SetupkeyHandler{
		redis:        r,
		config:       config,
		accountStore: account,
		serverStore:  server,
		user: user,
	}
}

type SetupKeyRequest struct {
	Group      *string
	Job        *string
	Permission *key.PermissionType
}

func (r SetupKeyRequest) IsValid() (bool, error) {
    if !(*r.Permission == key.RWXKey || *r.Permission == key.RWKey || *r.Permission == key.RKey || r.Permission != nil) {
        return false, fmt.Errorf("valid permission type")
    }

	if r.Group == nil || r.Job == nil {
        return false, fmt.Errorf("required parameter")
	}

    return true, nil
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

		_, err = req.IsValid()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sub := r.Header.Get("sub")
		setupKey, err := key.NewSetupKey(sub, *req.Group, *req.Job, *req.Permission)
		if err != nil {
			http.Error(w, "failed to create setupkey", http.StatusBadRequest)
			return
		}

		// TODO: create network
		// TODO: create group

		_, err = h.user.CreateUser(sub, "", *req.Group, *req.Permission)
		if err != nil {
			if errors.Is(err, model.ErrUserAlredyExists) {
				// TODO: get setup key, later return already exists setupkey
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, fmt.Sprintf("failed to create user. %v", err), http.StatusBadRequest)
			return
		}

		// TODO: create setup key

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(setupKey)
		return
	case http.MethodDelete:
		log.Println("delete setupkey")
		return
	default:
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}
