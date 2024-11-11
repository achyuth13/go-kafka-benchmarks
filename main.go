package main

import (
	"fmt"
	"go-kafka/producers"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Kafka struct {
		Brokers        []string `yaml:"brokers"`
		SegmentioTopic string   `yaml:"segmentiotopic"`
		ConfluentTopic string   `yaml:"confluenttopic"`
	} `yaml:"kafka"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	return &config, err
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize both producers
	segmentioProducer := producers.NewSegmentioProducer(config.Kafka.Brokers, config.Kafka.SegmentioTopic)
	defer segmentioProducer.Close()

	confluentProducer, err := producers.NewConfluentProducer(config.Kafka.Brokers[0], config.Kafka.ConfluentTopic)
	if err != nil {
		log.Fatalf("Error initializing Confluent producer: %v", err)
	}
	defer confluentProducer.Close()

	message := "Benchmarking message"
	start := time.Now()
	for i := 0; i < 1000; i++ {
		if err := segmentioProducer.Produce(message); err != nil {
			log.Printf("Segmentio producer error: %v", err)
		}
	}
	fmt.Printf("Segmentio Producer Time Taken: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < 1000; i++ {
		if err := confluentProducer.Produce(message); err != nil {
			log.Printf("Confluent producer error: %v", err)
		}
	}
	fmt.Printf("Confluent Producer Time Taken: %v\n", time.Since(start))
}
