package config

type Database struct {
	Address string
}

func NewDatabase(conf *Config) Database {
	db := Database{
		Address: conf.Address,
	}

	conf.DB = db

	return db
}
