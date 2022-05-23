package config

func New() (*Config, error) {
	return &Config{
		App: App{
			Port: "8080",
		},
	}, nil
}
