package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configName string = "dotsync.json"

type DotSyncConfig struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	IsDirectory bool   `json:"isDirectory"`
}

func ReadConfig() ([]DotSyncConfig, error) {
	data, err := os.ReadFile(configName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("dotsync.json config file not found")
		}
		return nil, err
	}

	var dotSyncConfig []DotSyncConfig
	if err = json.Unmarshal(data, &dotSyncConfig); err != nil {
		return nil, err
	}

	return dotSyncConfig, nil
}

func WriteConfig(dir []os.DirEntry) error {
	var dotSyncConfig []DotSyncConfig

	for _, entries := range dir {
		if entries.Name() == configName || entries.Name() == ".git" {
			continue
		}

		if entries.IsDir() {
			dotSyncConfig = append(dotSyncConfig, DotSyncConfig{
				Source:      entries.Name(),
				Target:      "",
				IsDirectory: true,
			})
		} else {
			dotSyncConfig = append(dotSyncConfig, DotSyncConfig{
				Source:      entries.Name(),
				Target:      "",
				IsDirectory: false,
			})
		}
	}

	jsonData, err := json.MarshalIndent(&dotSyncConfig, "", " ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(configName, jsonData, 0o644); err != nil {
		return err
	}

	return nil
}
