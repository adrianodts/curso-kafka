package main

import (
	"fmt"
	"log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChain := make(chan kafka.Event)
	producer := NewKafkaProducer()
	// envia para todas as partições quando o key não informado.
	// Publish("mensagem", "teste", producer, nil, deliveryChain)
	// envia sempre para a mesma partição quando o key for informado.
	Publish("mensagem", "teste", producer, []byte("transferencia"), deliveryChain)
	go AsyncDeliveryReport(deliveryChain)
	//SyncDeliveryReport(deliveryChain)
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "go-kafka_kafka_1:9092",
		"delivery.timeout.ms": "0",
		"acks": "all", // 0-Não aguarda confirmação de entrega; 1-Lider confirma a entrega: all- Lider e brokers confirmam a entrega
		"enable.idempotence": "true",
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
}

func AsyncDeliveryReport(deliveryChain chan kafka.Event) {
	for e := range deliveryChain {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar mensagem")
			} else {
				fmt.Println("Mensagem enviada", ev.TopicPartition)
			}
		}
	}
}

func SyncDeliveryReport(deliveryChain chan kafka.Event) {
	e := <-deliveryChain
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar mensagem")
	} else {
		fmt.Println("Mensagem enviada", msg.TopicPartition)
	}
}