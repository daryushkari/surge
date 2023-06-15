package config

import (
	"encoding/json"
	"os"
)

var (
	config *Config
)

type Config struct {
	ServiceName       string `json:"service-name"`
	Redis             string `json:"redis"`
	JaegerURL         string `json:"jaeger-url"`
	RequestTimeWindow int64  `json:"request-time-window"`
	RequestLiveTime   int64  `json:"request-live-time"`
	Database          struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DbName   string `json:"db-name"`
		Port     string `json:"port"`
	} `json:"database"`
	ExternalExpose struct {
		Rest string `json:"rest"`
	} `json:"external-expose"`
}

func InitCnf(cnfPath string) error {
	if config == nil {
		config = &Config{}
		file, err := os.ReadFile(cnfPath)
		if err != nil {
			return err
		}
		err = json.Unmarshal(file, &config)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCnf() *Config {
	return config
}
