package benchmarks_test

import (
	"go-kafka/config"
	"go-kafka/producers"
	"testing"
)

// Initialize benchmark settings, including config and producer setup
func BenchmarkSegmentioProducer(b *testing.B) {
	config, err := config.LoadConfig()
	if err != nil {
		b.Fatalf("Error loading config: %v", err)
	}

	segmentioProducer := producers.NewSegmentioProducer(config.Kafka.Brokers, config.Kafka.SegmentioTopic)
	defer segmentioProducer.Close()

	message := "Benchmarking message"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := segmentioProducer.Produce(message); err != nil {
			b.Errorf("Segmentio producer error: %v", err)
		}
	}
	b.StopTimer()
}
