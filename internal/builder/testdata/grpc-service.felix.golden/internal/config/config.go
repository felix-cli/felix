package config

// Config holds all of our vars set by the env
type Config struct {
	Host string `env:"HOST"`
	Port string `env:"PORT" envDefault:"3000"`
}

// New returns a new config struct
func New() *Config {
	return &Config{}
}
