package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/codephobia/pool-overlay/apps/api/pkg/overlay"
	"github.com/codephobia/pool-overlay/apps/api/pkg/state"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Server is an api server.
type Server struct {
	config  *Config
	db      *gorm.DB
	overlay *overlay.Overlay
	state   *state.State

	httpServer *http.Server
	router     *mux.Router
	version    map[string]*mux.Router
}

// NewServer returns a new api server.
func NewServer(config *Config, db *gorm.DB, overlay *overlay.Overlay, state *state.State) *Server {
	return &Server{
		config:  config,
		db:      db,
		overlay: overlay,
		state:   state,

		router:  mux.NewRouter().StrictSlash(true),
		version: make(map[string]*mux.Router),
	}
}

// Init initializes the api server.
func (server *Server) Init() {
	// versions
	server.AddVersion(VersionLatest)
	server.AddVersion(server.config.Version.Current)
	server.AddVersion(server.config.Version.Previous)

	// index
	server.AddRoute("/", server.handleIndex())

	// health check
	server.AddRoute("/health-check", server.handleHealthCheck())

	// Serve static files from /public
	// TODO: NEED A NICER WAY TO HANDLE ALL THESE SLASHES
	server.router.PathPrefix("/" + server.config.PublicDir + "/").Handler(
		http.StripPrefix(
			"/"+server.config.PublicDir+"/",
			http.FileServer(
				http.Dir("./"+server.config.PublicDir+"/"),
			),
		),
	)

	// init api routes
	server.InitRoutes()
}

// Run starts the api server.
func (server *Server) Run() error {
	hostURL := server.config.Host + ":" + server.config.Port

	// create the server
	server.httpServer = &http.Server{
		Handler:      server.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// create a listener
	listener, err := net.Listen("tcp", hostURL)
	if err != nil {
		return fmt.Errorf("error starting api server: %s", err)
	}

	// run server
	log.Printf("[INFO] API Server running: %s", listener.Addr().String())
	return server.httpServer.Serve(listener)
}

// Handler handles incoming api routes.
func (server *Server) Handler() http.Handler {
	return server.router
}
