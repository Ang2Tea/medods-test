package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

type StartMode string

const (
	DebugMode  StartMode = "debug"
	DockerMode StartMode = "docker"
	ProdMode   StartMode = "prod"
)

type Config struct {
	AppName   string    `yaml:"name"`
	Mode      StartMode `yaml:"mode"`
	WebServer struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"web-server"`
}

func GetConfig(configPath string) (c Config) {
	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		panic(err.Error())
	}

	err = yaml.Unmarshal(yamlData, &c)
	if err != nil {
		panic(err.Error())
	}

	return
}
