package main

import (
	"fmt"
	// kafka2 "github.com/xumes/fullcycle-emulator/application/kafka"
	"github.com/xumes/fullcycle-emulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "route.new-direction", producer)

	for msg := range msgChan {
		fmt.Println(string( msg.Value))
	}
	// route := route2.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()

	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
