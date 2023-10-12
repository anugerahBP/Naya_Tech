package config

type Config struct {
	Address string
	DB      Database
}

func NewConfig() Config {
	conf := Config{
		Address: "localhost:8080",
	}
	NewDatabase(&conf)

	return conf
}
