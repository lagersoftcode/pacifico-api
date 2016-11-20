package main

import "github.com/BurntSushi/toml"

type Config struct {
	ServerAddress string
	DBConnString  string
	AuthKey       string
	AppDomain     string
	CORSDomain    string
}

func GetConfig() Config {
	var config = Config{
		ServerAddress: ":8000",
		DBConnString:  "fer:password@/pacifico?charset=utf8&parseTime=True&loc=Local",
		AuthKey:       "123456",
		AppDomain:     "localhost",
		CORSDomain:    "*",
	}

	toml.DecodeFile("pacifico.conf", &config)
	return config
}
