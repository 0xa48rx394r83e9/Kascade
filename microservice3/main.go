// microservice3/main.go
package main

import (
    "log"
    "microservice3/config"
    "microservice3/consumer"
    "microservice3/producer"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Create Kafka producer
    producer := producer.NewEventProducer(cfg.KafkaBrokers)
    defer producer.Close()

    // Create Kafka consumer
    consumer := consumer.NewEventConsumer(cfg.KafkaBrokers, cfg.KafkaTopics)
    defer consumer.Close()

    // Start consuming events
    log.Println("Microservice3 started. Consuming events...")
    consumer.Consume()
}