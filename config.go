package main

import "github.com/BurntSushi/toml"

type Config struct {
	Address      string
	DBConnString string
	TokenKey     string
}

func GetConfig() Config {
	var config = Config{
		Address:      ":8000",
		DBConnString: "fer:password@/pacifico?charset=utf8&parseTime=True&loc=Local",
		TokenKey:     "123456",
	}

	toml.DecodeFile("pacifico.conf", &config)
	return config
}
