package config

import (
	"encoding/json"
	"os"
)

// loadTargets liest die Ziele aus einer JSON-Datei ein
func LoadTargets(filename string) ([]Target, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var targets []Target
	err = json.Unmarshal(data, &targets)
	if err != nil {
		return nil, err
	}

	return targets, nil
}
