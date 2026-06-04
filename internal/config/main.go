package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	filePath := filepath.Join(homeDir, ".gatorconfig.json")
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	return config, err
}
