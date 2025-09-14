package api

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	EnvConfig struct {
		ProxmoxHost     string `yaml:"proxmox_host"`
		ProxmoxPort     int    `yaml:"proxmox_port"`
		ProxmoxUsername string `yaml:"proxmox_username"`
		ProxmoxPassword string `yaml:"proxmox_password"`
	}

	AppConfig map[string]*EnvConfig
)

// Global configuration for the application.
var appConfig AppConfig
var EnvConfigProd EnvConfig

func LoadAppConfig() error {
	file, _ := os.Open("config.yml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&appConfig); err != nil {
		return err
	}

	return nil
}

func LoadEnvConfig(env string) EnvConfig {
	envConfig, ok := appConfig[env]
	if !ok {
		panic(fmt.Errorf("No such environment: %s", env))
	}

	return *envConfig
}

func PrintEnvConfig(config EnvConfig, env string) {
	fmt.Println("Environment:    " + env)
	fmt.Println("  Proxmox host:     ", config.ProxmoxHost)
	fmt.Println("  Proxmox port:     ", config.ProxmoxPort)
	fmt.Println("  Proxmox username: ", config.ProxmoxUsername)
	fmt.Println("  Proxmox password: ", config.ProxmoxPassword)

	fmt.Println()
}
