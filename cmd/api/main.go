package main

import (
	"StandartWebServer/internal/app/api"
	"log"
)

var ()

func init() {

}

func main() {
	log.Println("it working")
	//server instance init
	server := api.New()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
