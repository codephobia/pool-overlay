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

	// overlay/toggle/flags
	server.AddRouteToAllVersions("/overlay/toggle/flags", server.handleOverlayToggleFlags())

	// overlay/toggle/fargo
	server.AddRouteToAllVersions("/overlay/toggle/fargo", server.handleOverlayToggleFargo())

	// overlay/toggle/score
	server.AddRouteToAllVersions("/overlay/toggle/score", server.handleOverlayToggleScore())

	// web socket connection to telestrator
	server.AddRouteToAllVersions("/telestrator", server.handleTelestrator())

	// players
	server.AddRouteToAllVersions("/players", server.handlePlayers())

	// players/count
	server.AddRouteToAllVersions("/players/count", server.handlePlayersCount())

	// games
	server.AddRouteToAllVersions("/games", server.handleGames())

	// games/count
	server.AddRouteToAllVersions("/games/count", server.handleGamesCount())

	// players/{playerID}
	server.AddRouteToAllVersions("/players/{playerID}", server.handlePlayerByID())

	// flags
	server.AddRouteToAllVersions("/flags", server.handleFlags())

	// game
	server.AddRouteToAllVersions("/game", server.handleGame())

	// game/type
	server.AddRouteToAllVersions("/game/type", server.handleGameType())

	// game/vs-mode
	server.AddRouteToAllVersions("/game/vs-mode", server.handleGameVsMode())

	// game/race-to
	server.AddRouteToAllVersions("/game/race-to", server.handleGameRaceTo())

	// game/score
	server.AddRouteToAllVersions("/game/score", server.handleGameScore())

	// game/players
	server.AddRouteToAllVersions("/game/players", server.handleGamePlayers())

	// game/players/flag
	server.AddRouteToAllVersions("/game/players/flag", server.handleGamePlayersFlag())

	// game/players/name
	server.AddRouteToAllVersions("/game/players/name", server.handleGamePlayersName())

	// game/teams
	server.AddRouteToAllVersions("/game/teams", server.handleGameTeams())

	// game/fargo-hot-handicap
	server.AddRouteToAllVersions("/game/fargo-hot-handicap", server.handleGameFargoHotHandicap())
}
