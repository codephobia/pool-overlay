package api

import "net/http"

// AddRoute adds a route and handler to the router.
func (server *Server) AddRoute(path string, handler http.Handler) {
	server.router.Handle(path, handler)
}

// AddRouteToVersion adds a route to a versioned subrouter.
func (server *Server) AddRouteToVersion(version string, path string, handler http.Handler) {
	server.version[version].Handle(path, handler)

	// add current version routes to latest version
	if version == server.config.Version.Current {
		server.version[VersionLatest].Handle(path, handler)
	}
}

// AddRouteToAllVersions adds the route to all versions of the api.
func (server *Server) AddRouteToAllVersions(path string, handler http.Handler) {
	for v := range server.version {
		server.AddRouteToVersion(v, path, handler)
	}
}

// InitRoutes adds all routes to the api.
func (server *Server) InitRoutes() {
	// web socket connection to overlay
	server.AddRouteToAllVersions("/overlay", server.handleOverlay())

	// players
	server.AddRouteToAllVersions("/players", server.handlePlayers())

	// players/count
	server.AddRouteToAllVersions("/players/count", server.handlePlayersCount())
}
