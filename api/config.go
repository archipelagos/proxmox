package api

import (
	"os"
	"strconv"
)

type (
	EnvConfig struct {
		ProxmoxInternalAddress   string
		ProxmoxInternalTCPPort   int
		ProxmoxExternalAddress   string
		ProxmoxExternalTCPPort   int
		ProxmoxUsername          string
		ProxmoxPassword          string
		ProxmoxConnectionTimeout int
	}
)

// Global configuration for the application.
var EnvConfigProd EnvConfig

func LoadEnvConfig() EnvConfig {
	var envConfig EnvConfig
	var err error

	envConfig.ProxmoxInternalAddress = os.Getenv("PROXMOX_INTERNAL_ADDRESS_CICD_SECRET_VALUE")
	if len(envConfig.ProxmoxInternalAddress) == 0 {
		panic("Empty internal address")
	}

	envConfig.ProxmoxInternalTCPPort, err = strconv.Atoi(os.Getenv("PROXMOX_INTERNAL_TCP_PORT_CICD_SECRET_VALUE"))
	if err != nil {
		panic("Bad internal TCP port")
	}

	envConfig.ProxmoxExternalAddress = os.Getenv("PROXMOX_EXTERNAL_ADDRESS_CICD_SECRET_VALUE")
	if len(envConfig.ProxmoxInternalAddress) == 0 {
		panic("Empty external address")
	}

	envConfig.ProxmoxExternalTCPPort, err = strconv.Atoi(os.Getenv("PROXMOX_EXTERNAL_TCP_PORT_CICD_SECRET_VALUE"))
	if err != nil {
		panic("Bad external TCP port")
	}

	envConfig.ProxmoxUsername = os.Getenv("PROXMOX_USERNAME_CICD_SECRET_VALUE")
	if err != nil {
		panic("Empty username")
	}

	envConfig.ProxmoxPassword = os.Getenv("PROXMOX_PASSWORD_CICD_SECRET_VALUE")
	if err != nil {
		panic("Empty password")
	}

	envConfig.ProxmoxConnectionTimeout, err = strconv.Atoi(os.Getenv("PROXMOX_CONNECTION_TIMEOUT_CICD_SECRET_VALUE"))
	if err != nil || envConfig.ProxmoxConnectionTimeout == 0 {
		panic("Bad connection timeout")
	} else if envConfig.ProxmoxConnectionTimeout < 1 {
		panic("Negative connection timeout")
	}

	return envConfig
}
