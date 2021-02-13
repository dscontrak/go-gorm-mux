package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {	
	return &Config{
		DB: &DBConfig{
			Dialect:  "sqlite3",
			Username: "",
			Password: "",
			Name:     "./db.sqlite",
			Charset:  "utf8",
		},
	}
}
