package main

import (
	"fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChain := make(chan kafka.Event)
	producer := NewKafkaProducer()
	Publish("mensagem", "teste", producer, nil, deliveryChain)

	// Sincrono
	e := <-deliveryChain
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar mensagem")
	} else {
		fmt.Println("Mensagem enviada", msg.TopicPartition)
	}
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChain chan kafka.Event) error {
	message := &kafka.Message {
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key: key,
	}
	err := producer.Produce(message, deliveryChain)

	if err != nil {
		return err
	}
	return nil
	/*
	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "teste"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	*/
}