package challonge

// Config is the configuration required for Challonge.
type Config struct {
	APIKey   string
	Username string
}

// NewConfig returns a new Config.
func NewConfig(apiKey string, username string) *Config {
	return &Config{
		APIKey:   apiKey,
		Username: username,
	}
}
