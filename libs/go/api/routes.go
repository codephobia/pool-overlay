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
	server.AddRouteToAllVersions("/table/3/overlay/toggle", server.handleOverlayToggle(3))

	// overlay/toggle/flags
	// server.AddRouteToAllVersions("/overlay/toggle/flags", server.handleOverlayToggleFlags())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/flags", server.handleOverlayToggleFlags(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/flags", server.handleOverlayToggleFlags(2))
	server.AddRouteToAllVersions("/table/3/overlay/toggle/flags", server.handleOverlayToggleFlags(3))

	// overlay/toggle/fargo
	// server.AddRouteToAllVersions("/overlay/toggle/fargo", server.handleOverlayToggleFargo())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/fargo", server.handleOverlayToggleFargo(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/fargo", server.handleOverlayToggleFargo(2))
	server.AddRouteToAllVersions("/table/3/overlay/toggle/fargo", server.handleOverlayToggleFargo(3))

	// overlay/toggle/score
	// server.AddRouteToAllVersions("/overlay/toggle/score", server.handleOverlayToggleScore())
	server.AddRouteToAllVersions("/table/1/overlay/toggle/score", server.handleOverlayToggleScore(1))
	server.AddRouteToAllVersions("/table/2/overlay/toggle/score", server.handleOverlayToggleScore(2))
	server.AddRouteToAllVersions("/table/3/overlay/toggle/score", server.handleOverlayToggleScore(3))

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

	// table/swap
	server.AddRouteToAllVersions("/table/1/swap/{newTable}", server.handleTableSwap(1))
	server.AddRouteToAllVersions("/table/2/swap/{newTable}", server.handleTableSwap(2))
	server.AddRouteToAllVersions("/table/3/swap/{newTable}", server.handleTableSwap(3))

	// game
	server.AddRouteToAllVersions("/table/1/game", server.handleGame(1))
	server.AddRouteToAllVersions("/table/2/game", server.handleGame(2))
	server.AddRouteToAllVersions("/table/3/game", server.handleGame(3))

	// game/type
	server.AddRouteToAllVersions("/table/1/game/type", server.handleGameType(1))
	server.AddRouteToAllVersions("/table/2/game/type", server.handleGameType(2))
	server.AddRouteToAllVersions("/table/3/game/type", server.handleGameType(3))

	// game/vs-mode
	server.AddRouteToAllVersions("/table/1/game/vs-mode", server.handleGameVsMode(1))
	server.AddRouteToAllVersions("/table/2/game/vs-mode", server.handleGameVsMode(2))
	server.AddRouteToAllVersions("/table/3/game/vs-mode", server.handleGameVsMode(3))

	// game/race-to
	server.AddRouteToAllVersions("/table/1/game/race-to", server.handleGameRaceTo(1))
	server.AddRouteToAllVersions("/table/2/game/race-to", server.handleGameRaceTo(2))
	server.AddRouteToAllVersions("/table/3/game/race-to", server.handleGameRaceTo(3))

	// game/score
	server.AddRouteToAllVersions("/table/1/game/score", server.handleGameScore(1))
	server.AddRouteToAllVersions("/table/2/game/score", server.handleGameScore(2))
	server.AddRouteToAllVersions("/table/3/game/score", server.handleGameScore(3))

	// game/players
	server.AddRouteToAllVersions("/table/1/game/players", server.handleGamePlayers(1))
	server.AddRouteToAllVersions("/table/2/game/players", server.handleGamePlayers(2))
	server.AddRouteToAllVersions("/table/3/game/players", server.handleGamePlayers(3))

	// game/players/flag
	server.AddRouteToAllVersions("/game/players/flag", server.handleGamePlayersFlag())

	// game/players/name
	server.AddRouteToAllVersions("/game/players/name", server.handleGamePlayersName())

	// game/teams
	server.AddRouteToAllVersions("/game/teams", server.handleGameTeams())

	// game/fargo-hot-handicap
	server.AddRouteToAllVersions("/table/1/game/fargo-hot-handicap", server.handleGameFargoHotHandicap(1))
	server.AddRouteToAllVersions("/table/2/game/fargo-hot-handicap", server.handleGameFargoHotHandicap(2))
	server.AddRouteToAllVersions("/table/3/game/fargo-hot-handicap", server.handleGameFargoHotHandicap(3))

	// tournament
	server.AddRouteToAllVersions("/tournament", server.handleTournament())

	// tournament/list
	server.AddRouteToAllVersions("/tournament/list", server.handleTournamentList())

	// tournament/load
	server.AddRouteToAllVersions("/tournament/load", server.handleTournamentLoad())

	// tournament/unload
	server.AddRouteToAllVersions("/tournament/unload", server.handleTournamentUnload())

	// tournament/{tournamentID}
	server.AddRouteToAllVersions("/tournament/{tournamentID}", server.handleTournamentByID())
}
