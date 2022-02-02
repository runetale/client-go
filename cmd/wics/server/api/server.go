package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/api/handler"
	"github.com/Notch-Technologies/wizy/store"
)

func newMuxHandler(
	config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, r *redis.RedisClient, user *redis.UserStore,
	network *redis.NetworkStore, group *redis.OrgGroupStore,
) *http.ServeMux {
	mux := http.NewServeMux()

	sh := handler.NewSetupKeyHanlder(r, config, account, server, user, network, group)
	// admin
	mux.Handle("/api/setupkey", adminMiddleware(http.HandlerFunc(sh.SetupKey)))

	// manager

	// default

	return mux
}

func newApiServer(mux *http.ServeMux, port uint16) *http.Server {
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

func NewHTTPServer(
	port uint16, config *config.Config,
	account *redis.AccountStore, server *store.ServerStore,
	user *redis.UserStore, r *redis.RedisClient, 
	network *redis.NetworkStore, group *redis.OrgGroupStore,
) *http.Server {
	mux := newMuxHandler(config, account, server, r, user, network, group)
	return newApiServer(mux, port)
}
