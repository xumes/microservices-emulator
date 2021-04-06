package main

import (
	"fmt"
	route2 "github.com/xumes/fullcycle-emulador/application/route"
)

func main() {
	route := route2.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
