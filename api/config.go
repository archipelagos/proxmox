package api

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	EnvConfig struct {
		ProxmoxInternalAddress string `yaml:"proxmox_internal_address"`
		ProxmoxInternalTCPPort int    `yaml:"proxmox_internal_tcp_port"`
		ProxmoxExternalAddress string `yaml:"proxmox_external_address"`
		ProxmoxExternalTCPPort int    `yaml:"proxmox_external_tcp_port"`
		ProxmoxUsername        string `yaml:"proxmox_username"`
		ProxmoxPassword        string `yaml:"proxmox_password"`
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
	fmt.Println("Environment:                 " + env)
	fmt.Println("  Proxmox internal address:  ", config.ProxmoxInternalAddress)
	fmt.Println("  Proxmox internal TCP port: ", config.ProxmoxInternalTCPPort)
	fmt.Println("  Proxmox external address:  ", config.ProxmoxExternalAddress)
	fmt.Println("  Proxmox external TCP port: ", config.ProxmoxExternalTCPPort)
	fmt.Println("  Proxmox username:          ", config.ProxmoxUsername)
	fmt.Println("  Proxmox password:          ", config.ProxmoxPassword)

	fmt.Println()
}
