package main

import (
	"StandartWebServer/internal/app/api"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string = "configs/api.toml"
)

func init() {
	// На этапе запуска получем путь к конфигу
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml")
}

func main() {
	flag.Parse()
	log.Println("it working")
	//server instance init
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("can not find configs file", err)
	}

	server := api.New(config)

	log.Fatal(server.Start())
}
