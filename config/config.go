package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

//Config ...
type Config struct {
	Application `json:"application"`
	MariaDB     `json:"mariadb"`
	Redis       `json:"redis"`
}

func LoadConfiguration(file string) Config {
	var config Config
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile, err := os.Open(filepath.Join(workDir, file))
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		panic(err)
	}
	return config
}
