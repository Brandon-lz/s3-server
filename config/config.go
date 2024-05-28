package config

import (
	"go_learn/utils"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var Config ConfigDF

// [service]
// deploy_host = "192.168.70.213:30577"
// port = 8060

// [logging]
// level = "debug"

// [redis]
// host = "localhost"
// port = 6379
// db = 0
// password = ""

type ConfigDF struct {
	Service struct {
		DeployHost string `json:"deploy_host"`
		Port       int    `json:"port"`
	} `json:"service"`
	Logging struct {
		Level string `json:"level"`
	} `json:"logging"`
	Redis struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		DB       int    `json:"db"`
		Password string `json:"password"`
	} `json:"redis"`
}

func LoadConfig() error {
	tomlFilePath := ""
	if os.Getenv("deploy_env") == "" {
		log.Fatalf("deploy_env not set")
	}
	tomlFilePath = "config-" + os.Getenv("deploy_env") + ".toml"
	var tomlData map[string]interface{}
	if _, err := toml.DecodeFile(tomlFilePath, &tomlData); err != nil {
		return err
	}
	utils.DeserializeData(tomlData, &Config)
	return nil
}
