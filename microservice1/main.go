// microservice1/main.go
package main

import (
    "log"
    "microservice1/config"
    "microservice1/consumer"
    "microservice1/producer"
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
    log.Println("Microservice1 started. Consuming events...")
    consumer.Consume()
}