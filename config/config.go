// config/config.go
package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Kafka struct {
		Brokers        []string `yaml:"brokers"`
		SegmentioTopic string   `yaml:"segmentiotopic"`
		ConfluentTopic string   `yaml:"confluenttopic"`
	} `yaml:"kafka"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("./config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	return &config, err
}
