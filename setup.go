package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amimof/huego"
	"gopkg.in/yaml.v3"
)

func getConfigFileName() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.lights.yaml", homeDir)
}

func writeConfig(config Config) {
	configFile := getConfigFileName()
	data, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(configFile, data, 0644)
}

func readConfig() Config {
	config := Config{}
	configFile, err := os.ReadFile(getConfigFileName())
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(configFile, &config)
	return config
}

func getBridge() huego.Bridge {
	config := readConfig()
	bridge := huego.New(config.IpAddress, config.ApiKey)
	if bridge == nil {
		log.Fatal("Failed to initialize bridge. If this continues, run setup again.")
	}
	return *bridge
}

func Setup() {
	config := Config{}
	bridge, err := huego.Discover()
	if err != nil {
		log.Fatal(err)
	}
	config.IpAddress = bridge.Host

	fmt.Println("Go press the button on your Hue bridge. Once you've done that, press Enter")
	fmt.Scanln()
	user, err := bridge.CreateUser("lights cli")
	if err != nil {
		log.Fatal(err)
	}
	config.ApiKey = user
	writeConfig(config)
	fmt.Println("Setup complete")
}
