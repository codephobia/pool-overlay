package api

// Config is the configuration for the API package.
type Config struct {
	Host      string   `json:"host"`
	Port      string   `json:"port"`
	PublicDir string   `json:"public_dir"`
	Version   *Version `json:"version"`
}
