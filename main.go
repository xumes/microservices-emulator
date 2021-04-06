package main

import (
	"fmt"
	route2 "github.com/xumes/fullcycle-emulador/application/route"
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
	route := route2.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
