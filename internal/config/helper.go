package config

import (
	"encoding/json"
	"os"
)

func write(c Config) error {
	jsonData, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileName, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}
