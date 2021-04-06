package main

import (
	// "fmt"
	// kafka2 "github.com/xumes/fullcycle-emulator/application/kafka"
	"github.com/xumes/fullcycle-emulador/infra/kafka"
	// ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
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

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "route.new-direction", producer)

	for {
		_ = 1
	}
	// route := route2.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()

	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}
