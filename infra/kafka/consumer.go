package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"log"
	"fmt"
)

type KafkaConsomer struct {
	MsgChan chan *ckafka.Messsage
}


func(k *KafkaConsomer) Consume() {
	configMap := &ckafka.configMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id": os.Getenv("KafkaConsumerGroupdId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{os.Getenv("kafkaReadTopic")}
	c.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage( -1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
