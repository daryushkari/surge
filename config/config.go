package config

import (
	"encoding/json"
	"os"
)

var (
	config *Config
)

type Config struct {
	ServiceName string `json:"service-name"`
	Redis       string `json:"redis"`
	JaegerURL   string `json:"jaeger-url"`
	// request-time-window is 10 minutes in milliseconds
	RequestTimeWindow       int64    `json:"request-time-window"`
	RequestLiveTime         int64    `json:"request-live-time"`
	NotifyPricingRetryCount int      `json:"notify-pricing-retry-count"`
	OverpassUrl             string   `json:"overpass-url"`
	NotifyPricingTimeout    int      `json:"notify-pricing-timeout"`
	Nats                    []string `json:"nats"`
	PricingSubject          string   `json:"pricing-subject"`
	Database                struct {
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
