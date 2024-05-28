// microservice3/consumer/event_consumer.go
package consumer

import (
    "encoding/json"
    "log"
    "microservice3/handler"
    "microservice3/model"

    "github.com/Shopify/sarama"
)

type EventConsumer struct {
    consumer sarama.Consumer
    topics   []string
    handler  *handler.EventHandler
}

func NewEventConsumer(brokers []string, topics []string) *EventConsumer {
    consumer, err := sarama.NewConsumer(brokers, nil)
    if err != nil {
        log.Fatalf("Failed to create Kafka consumer: %v", err)
    }

    return &EventConsumer{
        consumer: consumer,
        topics:   topics,
        handler:  handler.NewEventHandler(),
    }
}

func (c *EventConsumer) Close() {
    c.consumer.Close()
}

func (c *EventConsumer) Consume() {
    for _, topic := range c.topics {
        partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
        if err != nil {
            log.Printf("Failed to consume events from topic %s: %v", topic, err)
            continue
        }

        go func(pc sarama.PartitionConsumer) {
            defer pc.Close()

            for msg := range pc.Messages() {
                var event model.Event
                err := json.Unmarshal(msg.Value, &event)
                if err != nil {
                    log.Printf("Failed to unmarshal event: %v", err)
                    continue
                }

                c.handler.HandleEvent(&event)
            }
        }(partitionConsumer)
    }

    // Wait indefinitely
    select {}
}