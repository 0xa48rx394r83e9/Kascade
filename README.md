# Kascade

This project demonstrates a distributed, event-driven microservices architecture using Kafka and Go. It consists of three microservices that communicate with each other through Kafka topics, allowing for loose coupling and scalability.

## Project Structure

The project is structured as follows:

```
project/
├── docker-compose.yml
├── README.md
├── microservice1/
│ ├── Dockerfile
│ ├── main.go
│ ├── handler/
│ │ └── event_handler.go
│ ├── producer/
│ │ └── event_producer.go
│ ├── consumer/
│ │ └── event_consumer.go
│ ├── model/
│ │ └── event.go
│ └── config/
│ └── config.go
├── microservice2/
│ ├── ...
├── microservice3/
│ ├── ...
└── kubernetes/
├── kafka/
│ ├── kafka-deployment.yaml
│ └── kafka-service.yaml
├── microservice1/
│ ├── deployment.yaml
│ └── service.yaml
├── microservice2/
│ ├── ...
└── microservice3/
├── ...

```

## Prerequisites

Before running the project, ensure that you have the following prerequisites installed:

- Docker
- Docker Compose
- Go (version 1.16 or higher)
- Kubernetes (optional, for deployment)

## Getting Started

To get started with the project, follow these steps:

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/your-project.git
   cd your-project
   ```

2. Build and run the services using Docker Compose:

   ```shell
   docker-compose up --build
   ```

   This command will build the Docker images for the microservices and start all the services defined in the `docker-compose.yml` file.

3. The microservices will start consuming and processing events from the respective Kafka topics.

## Microservices

The project includes three microservices:

- `microservice1`: Consumes events from `topic1` and `topic2`, processes them, and performs specific actions.
- `microservice2`: Consumes events from `topic2` and `topic3`, processes them, and performs specific actions.
- `microservice3`: Consumes events from `topic1` and `topic3`, processes them, and performs specific actions.

Each microservice is implemented in Go and follows a similar structure:

- `main.go`: The entry point of the microservice, where the Kafka producer and consumer are initialized, and the event consumption is started.
- `handler/`: Contains the event handler implementation for processing consumed events.
- `producer/`: Contains the Kafka producer implementation for producing events to Kafka topics.
- `consumer/`: Contains the Kafka consumer implementation for consuming events from Kafka topics.
- `model/`: Defines the data models and structs used by the microservice.
- `config/`: Contains the configuration loading and management code.

## Deployment

The project includes Kubernetes manifest files for deploying the microservices and Kafka to a Kubernetes cluster. To deploy the services, follow these steps:

1. Ensure you have a running Kubernetes cluster and `kubectl` configured to access it.

2. Apply the Kubernetes manifests:

   ```shell
   kubectl apply -f kubernetes/kafka/
   kubectl apply -f kubernetes/microservice1/
   kubectl apply -f kubernetes/microservice2/
   kubectl apply -f kubernetes/microservice3/
   ```

   This will create the necessary deployments and services for Kafka and the microservices in your Kubernetes cluster.

## Configuration

The microservices can be configured using environment variables. The available environment variables are:

- `KAFKA_BROKERS`: Comma-separated list of Kafka broker addresses (default: `kafka:9092`).
- `KAFKA_TOPICS`: Comma-separated list of Kafka topics the microservice consumes from.

You can modify the `docker-compose.yml` file or the Kubernetes deployment manifests to set the desired configuration values.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
