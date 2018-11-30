package config

import (
	"encoding/json"
  "os"
)

type Config struct {
  Port string `json:"port"`
}

var ServerConfig Config

func Init() {
  file := "config/config_local.json"
  configFile, _ := os.Open(file)
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&ServerConfig)
}
