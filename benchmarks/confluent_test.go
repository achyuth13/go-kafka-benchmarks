package benchmarks_test

import (
	"go-kafka/config"
	"go-kafka/producers"
	"testing"
)

// Initialize benchmark settings, including config and producer setup
func BenchmarkConfluentProducer(b *testing.B) {
	config, err := config.LoadConfig()
	if err != nil {
		b.Fatalf("Error loading config: %v", err)
	}

	confluentProducer, err := producers.NewConfluentProducer(config.Kafka.Brokers[0], config.Kafka.ConfluentTopic)
	if err != nil {
		b.Fatalf("Error initializing Confluent producer: %v", err)
	}
	defer confluentProducer.Close()

	message := "Benchmarking message"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := confluentProducer.Produce(message); err != nil {
			b.Errorf("Confluent producer error: %v", err)
		}
	}
	b.StopTimer()
}
