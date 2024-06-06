package server

type Config struct {
	port int
}

func NewConfig(opts ...func(config *Config)) *Config {
	config := &Config{
		port: 9090,
	}
	for _, opt := range opts {
		opt(config)
	}
	return config
}

func WithPort(port int) func(config *Config) {
	return func(config *Config) {
		config.port = port
	}
}
