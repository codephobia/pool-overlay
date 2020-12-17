package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// API struct
type API struct {
	config  *Config
	overlay *overlay.Overlay
	server  *http.Server
}

// NewAPI creates a new API struct.
func NewAPI(c *Config, overlay *overlay.Overlay) *API {
	return &API{
		config:  c,
		overlay: overlay,
	}
}

// Init initializes the API.
func (api *API) Init() error {
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	// create the server
	api.server = &http.Server{
		Handler:      handlers.CORS(allowedHeaders, allowedOrigins)(api.Handler()),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// create a listener
	hostURL := fmt.Sprintf("%s:%s", api.config.Host, api.config.Port)
	listener, err := net.Listen("tcp", hostURL)
	if err != nil {
		return fmt.Errorf("[ERROR] unable to start api server: %s", err)
	}

	// run server
	log.Printf("[INFO] API Server running: %s", listener.Addr().String())
	go api.server.Serve(listener)

	return nil
}

// Handler returns a mux router.
func (api *API) Handler() http.Handler {
	// create router
	r := mux.NewRouter().StrictSlash(true)

	// TODO: NEED A NICER WAY TO HANDLE ALL THESE SLASHES
	// Serve static files from /public
	r.PathPrefix("/" + api.config.PublicDir + "/").Handler(
		http.StripPrefix(
			"/"+api.config.PublicDir+"/",
			http.FileServer(
				http.Dir("./"+api.config.PublicDir+"/"),
			),
		),
	)

	// handle web sockets
	r.Handle("/overlay", api.handleOverlay())

	// return router
	return r
}
