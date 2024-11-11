package producers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConfluentProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewConfluentProducer(brokers string, topic string) (*ConfluentProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return nil, err
	}
	return &ConfluentProducer{Producer: producer, Topic: topic}, nil
}

func (p *ConfluentProducer) Produce(message string) error {
	return p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
}

func (p *ConfluentProducer) Close() {
	p.Producer.Close()
}
