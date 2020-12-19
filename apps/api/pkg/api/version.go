package api

// VersionLatest is the latest version for the API.
// We should only be using this for development.
var VersionLatest = "latest"

// Version is loaded from the config and used to allow for one version back.
type Version struct {
	Current  string `json:"current"`
	Previous string `json:"previous"`
}

// AddVersion adds a new sub router with version prefix
func (server *Server) AddVersion(version string) {
	var prefix string

	// Latest endpoint doesn't use the "v" prefix.
	if version == VersionLatest {
		prefix = "/" + version
	} else {
		prefix = "/v" + version
	}

	// Add subrouter for version being added.
	server.version[version] = server.router.PathPrefix(prefix).Subrouter()
}
