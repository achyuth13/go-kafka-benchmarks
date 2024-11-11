package producers

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type SegmentioProducer struct {
	Writer *kafka.Writer
}

func NewSegmentioProducer(brokers []string, topic string) *SegmentioProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Async:        false,
		RequiredAcks: int(kafka.RequireAll),
	})
	return &SegmentioProducer{Writer: writer}
}

func (p *SegmentioProducer) Produce(message string) error {
	return p.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key"),
			Value: []byte(message),
			Time:  time.Now(),
		},
	)
}

func (p *SegmentioProducer) Close() {
	p.Writer.Close()
}
