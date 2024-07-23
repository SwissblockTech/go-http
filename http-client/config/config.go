package config

import (
	"github.com/swisslockTech/go-http/http-client/logging"
	"github.com/swisslockTech/go-http/http-client/utils"
)

const (
	enableMonitoringEnvVar  = "ENABLE_MONITORING" // bool
	enableMonitoringDefault = true
)

func LoadConfig() *Config {
	logging.Log.Info("Load global configurations")

	return &Config{
		EnableMonitoring: utils.GetBoolEnv(enableMonitoringEnvVar, enableMonitoringDefault),
	}
}
