// microservice3/producer/event_producer.go
package producer

import (
    "encoding/json"
    "log"
    "microservice3/model"

    "github.com/Shopify/sarama"
)

type EventProducer struct {
    producer sarama.SyncProducer
}

func NewEventProducer(brokers []string) *EventProducer {
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true

    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        log.Fatalf("Failed to create Kafka producer: %v", err)
    }

    return &EventProducer{
        producer: producer,
    }
}

func (p *EventProducer) Close() {
    p.producer.Close()
}

func (p *EventProducer) ProduceEvent(topic string, event *model.Event) error {
    jsonEvent, err := json.Marshal(event)
    if err != nil {
        return err
    }

    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.ByteEncoder(jsonEvent),
    }

    _, _, err = p.producer.SendMessage(msg)
    return err
}