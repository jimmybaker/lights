package main

// Config ...
type Config struct {
	IpAddress string `yaml:"ipAddress"`
	ApiKey    string `yaml:"apiKey"`
}

// HueBridgeConnectionInfo ...
type HueBridgeConnectionInfo struct {
	Name      string
	IpAddress string
}
