package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupkeyHandlerManager interface {
	SetupKey(w http.ResponseWriter, r *http.Request)
}

type SetupkeyHandler struct {
	db              *database.Sqlite
	setupKeyUsecase *usecase.SetupKeyUsecase
}

func NewSetupKeyHanlder(
	db *database.Sqlite,
) *SetupkeyHandler {
	su := usecase.NewSetupKeyUsecase(db)
	return &SetupkeyHandler{
		setupKeyUsecase: su,
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
		//h.setupKeyUsecase.CreateSetupKey()
		//
		//sub := r.Header.Get("sub")
		//setupKey, err := h.setupKeyRepository.CreateSetupKey(sub, *req.Group, *req.Job, *req.Network, *req.Permission)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}

		w.Header().Set("Content-Type", "application/json")
		//json.NewEncoder(w).Encode(setupKey)
		return
	case http.MethodDelete:
		log.Println("delete setupkey")
		return
	default:
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}
