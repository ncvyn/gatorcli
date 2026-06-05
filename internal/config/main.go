package config

import (
	"encoding/json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName string = ".gatorconfig.json"

func Read() (Config, error) {
	fileContent, err := read(configFileName)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	return config, err
}
