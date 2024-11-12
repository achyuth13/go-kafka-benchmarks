package main

import (
	"fmt"
	"go-kafka/config"
	"go-kafka/producers"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func profileConfluentProducer(config *config.Config) {
	confluentProducer, err := producers.NewConfluentProducer(config.Kafka.Brokers[0], config.Kafka.ConfluentTopic)
	if err != nil {
		log.Fatalf("Error initializing Confluent producer: %v", err)
	}
	defer confluentProducer.Close()

	message := "Benchmarking message"
	start := time.Now()
	for i := 0; i < 1; i++ {
		if err := confluentProducer.Produce(message); err != nil {
			log.Printf("Confluent producer error: %v", err)
		}
	}
	fmt.Printf("Confluent Producer Time Taken: %v\n", time.Since(start))
}

func profileSegmentioProducer(config *config.Config) {
	segmentioProducer := producers.NewSegmentioProducer(config.Kafka.Brokers, config.Kafka.SegmentioTopic)
	defer segmentioProducer.Close()

	message := "Benchmarking message"
	start := time.Now()
	for i := 0; i < 1; i++ {
		if err := segmentioProducer.Produce(message); err != nil {
			log.Printf("Segmentio producer error: %v", err)
		}
	}
	fmt.Printf("Segmentio Producer Time Taken: %v\n", time.Since(start))
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil)) // Starts the pprof server
	}()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// First, profile Confluent producer
	fmt.Println("Starting Confluent producer profiling...")
	profileConfluentProducer(config)

	// Then, profile Segmentio producer
	fmt.Println("Starting Segmentio producer profiling...")
	profileSegmentioProducer(config)

	select {}
}
