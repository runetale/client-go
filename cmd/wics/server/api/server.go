package api

import (
	"log"
	"net/http"
	"strconv"
)

func newMuxHandler() *http.ServeMux {
	mux := http.NewServeMux()

	// admin
	mux.Handle("/api/setupkey", adminMiddleware(http.HandlerFunc(setupKey)))

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

func NewHTTPServer(port uint16) *http.Server {
	mux := newMuxHandler()
	return newApiServer(mux, port)
}
