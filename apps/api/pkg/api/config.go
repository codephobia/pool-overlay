package api

// Config is the configuration for the API package.
type Config struct {
	Scheme    string `json:"scheme"`
	Host      string `json:"host"`
	Port      string `json:"port"`
	PublicDir string `json:"public_dir"`
}
