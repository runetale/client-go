package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/redis"
	"github.com/Notch-Technologies/wizy/cmd/server/grpc/api/handler"
	"github.com/Notch-Technologies/wizy/store"
)

func newMuxHandler(
	config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, r *redis.RedisClient, user *redis.UserStore,
	network *redis.NetworkStore, group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) *http.ServeMux {
	mux := http.NewServeMux()

	sh := handler.NewSetupKeyHanlder(r, config, account, server, user, network, group, setupKey)
	// admin
	mux.Handle("/api/setupkey", adminMiddleware(http.HandlerFunc(sh.SetupKey)))

	// manager

	// default

	return mux
}

func listen(mux *http.ServeMux, port uint16) *http.Server {
	httpsrv := &http.Server{
		Addr:    ":" + strconv.Itoa(int(port)),
		Handler: mux,
	}

	go func() {
		if err := httpsrv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	return httpsrv
}

func NewAPIServer(
	port uint16, config *config.Config,
	account *redis.AccountStore, server *store.ServerStore,
	user *redis.UserStore, r *redis.RedisClient, 
	network *redis.NetworkStore, group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) *http.Server {
	mux := newMuxHandler(config, account, server, r, user, network, group, setupKey)
	return listen(mux, port)
}
