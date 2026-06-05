package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func read(fileName string) ([]byte, error) {
	filePath, err := getHomeFilePath(fileName)
	if err != nil {
		return nil, err
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func (c Config) write(fileName string) error {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath, err := getHomeFilePath(fileName)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func getHomeFilePath(fileName string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, fileName)
	return filePath, nil
}
