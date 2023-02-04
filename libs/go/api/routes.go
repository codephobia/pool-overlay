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
	// server.AddRouteToAllVersions("/overlay/toggle", server.handleOverlayToggle())
	server.AddRouteToAllVersions("/table/1/overlay/toggle", server.handleOverlayToggle(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle", server.handleOverlayToggle(2))

	// overlay/toggle/flags
	// server.AddRouteToAllVersions("/overlay/toggle/flags", server.handleOverlayToggleFlags())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/flags", server.handleOverlayToggleFlags(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/flags", server.handleOverlayToggleFlags(2))

	// overlay/toggle/fargo
	// server.AddRouteToAllVersions("/overlay/toggle/fargo", server.handleOverlayToggleFargo())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/fargo", server.handleOverlayToggleFargo(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/fargo", server.handleOverlayToggleFargo(2))

	// overlay/toggle/score
	// server.AddRouteToAllVersions("/overlay/toggle/score", server.handleOverlayToggleScore())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/score", server.handleOverlayToggleScore(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/score", server.handleOverlayToggleScore(2))

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
	// server.AddRouteToAllVersions("/game", server.handleGame())
	server.AddRouteToAllVersions("/table/1/game", server.handleGame(1))
	server.AddRouteToAllVersions("/table/2/game", server.handleGame(2))

	// game/type
	// server.AddRouteToAllVersions("/game/type", server.handleGameType())
	server.AddRouteToAllVersions("/table/1/game/type", server.handleGameType(1))
	server.AddRouteToAllVersions("/table/2/game/type", server.handleGameType(2))

	// game/vs-mode
	// server.AddRouteToAllVersions("/game/vs-mode", server.handleGameVsMode())
	server.AddRouteToAllVersions("/table/1/game/vs-mode", server.handleGameVsMode(1))
	server.AddRouteToAllVersions("/table/2/game/vs-mode", server.handleGameVsMode(2))

	// game/race-to
	// server.AddRouteToAllVersions("/game/race-to", server.handleGameRaceTo())
	server.AddRouteToAllVersions("/table/1/game/race-to", server.handleGameRaceTo(1))
	server.AddRouteToAllVersions("/table/2/game/race-to", server.handleGameRaceTo(2))

	// game/score
	// server.AddRouteToAllVersions("/game/score", server.handleGameScore())
	server.AddRouteToAllVersions("/table/1/game/score", server.handleGameScore(1))
	server.AddRouteToAllVersions("/table/2/game/score", server.handleGameScore(2))

	// game/players
	// server.AddRouteToAllVersions("/game/players", server.handleGamePlayers())
	server.AddRouteToAllVersions("/table/1/game/players", server.handleGamePlayers(1))
	server.AddRouteToAllVersions("/table/2/game/players", server.handleGamePlayers(2))

	// game/players/flag
	server.AddRouteToAllVersions("/game/players/flag", server.handleGamePlayersFlag())

	// game/players/name
	server.AddRouteToAllVersions("/game/players/name", server.handleGamePlayersName())

	// game/teams
	server.AddRouteToAllVersions("/game/teams", server.handleGameTeams())

	// game/fargo-hot-handicap
	// server.AddRouteToAllVersions("/game/fargo-hot-handicap", server.handleGameFargoHotHandicap())
	server.AddRouteToAllVersions("/table/1/game/fargo-hot-handicap", server.handleGameFargoHotHandicap(1))
	server.AddRouteToAllVersions("/table/2/game/fargo-hot-handicap", server.handleGameFargoHotHandicap(2))
}
