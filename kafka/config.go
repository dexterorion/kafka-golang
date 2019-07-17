package kafka

import (
	"time"

	kafka "github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// Configure configure.
func Configure(kafkaBrokerUrls []string, clientID string, topic string) (w *kafka.Writer, err error) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientID,
	}

	config := kafka.WriterConfig{
		Brokers:      kafkaBrokerUrls,
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		Dialer:       dialer,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		// CompressionCodec: snappy.NewCompressionCodec(),
	}
	w = kafka.NewWriter(config)
	writer = w
	return w, nil
}
