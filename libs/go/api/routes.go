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

	// overlay/toggle
	server.AddRouteToAllVersions("/overlay/toggle", server.handleOverlayToggle())

	// players
	server.AddRouteToAllVersions("/players", server.handlePlayers())

	// players/count
	server.AddRouteToAllVersions("/players/count", server.handlePlayersCount())

	// players/{playerID}
	server.AddRouteToAllVersions("/players/{playerID}", server.handlePlayerByID())

	// game
	server.AddRouteToAllVersions("/game", server.handleGame())

	// game/update/type
	server.AddRouteToAllVersions("/game/update/type", server.handleGameType())

	// game/update/vs-mode
	server.AddRouteToAllVersions("/game/update/vs-mode", server.handleGameVsMode())

	// game/update/race-to
	server.AddRouteToAllVersions("/game/update/race-to", server.handleGameRaceTo())

	// game/update/score
	server.AddRouteToAllVersions("/game/update/score", server.handleGameScore())

	// game/update/score/reset
	server.AddRouteToAllVersions("/game/update/score/reset", server.handleGameScoreReset())

	// game/update/players
	server.AddRouteToAllVersions("/game/update/players", server.handleGamePlayers())

	// game/update/players/flag
	server.AddRouteToAllVersions("/game/update/players/flag", server.handleGamePlayersFlag())

	// game/update/players/name
	server.AddRouteToAllVersions("/game/update/players/name", server.handleGamePlayersName())

	// TODO: IMPLEMENT THIS
	// game/update/players/save
	// server.AddRouteToAllVersions("/game/update/players/save", server.handleGamePlayers())

	// game/update/players/unset
	server.AddRouteToAllVersions("/game/update/players/unset", server.handleGamePlayersUnset())

	// game/update/teams
	server.AddRouteToAllVersions("/game/update/teams", server.handleGameTeams())
}
