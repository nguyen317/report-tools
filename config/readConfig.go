package config

import "github.com/jinzhu/configor"

type Config struct {
	App struct {
		Keyapp   string
		Token    string
		Idboard  string
		Username string
	}
	Server struct {
		Host string
		Port string
	}
	Database struct {
		Name     string
		Username string
		Password string
	}

	Contacts struct {
		Email string
	}
}

func ReadConfig() (cg Config) {
	configor.Load(&cg, "config.yml")
	return
}
