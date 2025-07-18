package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Fullscreen     bool   `yaml:"full_screen"`
	ScreenWidth    int    `yaml:"screen_width"`
	ScreenHeight   int    `yaml:"screen_height"`
	SpeedInSeconds int    `yaml:"speed_in_seconds"`
	Effect         string `yaml:"effect"`
}

func LoadConfig() *Config {
	result := Config{}
	b, err := os.ReadFile("./config.yml")

	if err != nil {
		return &Config{
			Fullscreen:     false,
			ScreenWidth:    0,
			ScreenHeight:   0,
			SpeedInSeconds: 5,
			Effect:         "crossfade",
		}
	}

	err = yaml.Unmarshal(b, &result)

	if err != nil {
		log.Fatalf("Error unmarshalling config file: %v", err)
	}

	if result.Effect == "" {
		result.Effect = "fade"
	}

	return &result
}
