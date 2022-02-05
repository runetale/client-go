package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupkeyHandlerManager interface {
	SetupKey(w http.ResponseWriter, r *http.Request)
}

type SetupkeyHandler struct {
	redis              *redis.RedisClient
	config             *config.Config
	accountStore       *redis.AccountStore
	serverStore        *store.ServerStore
	userStore          *redis.UserStore
	networkStore       *redis.NetworkStore
	orgGroupStore      *redis.OrgGroupStore
	setupKeyStore      *redis.SetupKeyStore
	setupKeyRepository *repository.SetupKeyRepository
}

func NewSetupKeyHanlder(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, user *redis.UserStore, network *redis.NetworkStore,
	group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) *SetupkeyHandler {
	setupKeyRepository := repository.NewSetupKeyRepository(r, config, account, server, user, network, group, setupKey)

	return &SetupkeyHandler{
		redis:              r,
		config:             config,
		accountStore:       account,
		serverStore:        server,
		userStore:          user,
		networkStore:       network,
		orgGroupStore:      group,
		setupKeyStore:      setupKey,
		setupKeyRepository: setupKeyRepository,
	}
}

type SetupKeyRequest struct {
	Group      *string
	Network    *string
	Job        *string
	Permission *key.PermissionType
}

func (r SetupKeyRequest) IsValid() (bool, error) {
	if !(*r.Permission == key.RWXKey || *r.Permission == key.RWKey || *r.Permission == key.RKey || r.Permission != nil) {
		return false, fmt.Errorf("valid permission type")
	}

	if r.Group == nil || r.Job == nil || r.Network == nil {
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
		setupKey, err := h.setupKeyRepository.CreateSetupKey(sub, *req.Group, *req.Job, *req.Network, *req.Permission)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
