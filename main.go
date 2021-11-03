package main

import (
	"flag"
	"log"

	"github.com/Dmitry-44/telegramBot/config"
	"github.com/Dmitry-44/telegramBot/connections"
)

func main() {
	
	path := flag.String(`config-path`, `config.dist.yaml`, `path to file config`)
	flag.Parse()

	log.Println(`start service "static-image"`)
	defer func() {
		log.Println(`stop service "static-image"`)
	}()

	conf, err := config.LoadConfig(*path)
	if err != nil {
		log.Fatalf(`load config: %s`, err.Error())
	}

	err = connections.Init(conf.connections)
	if err != nil {
		log.Fatalf(`init connections: %s`, err.Error())
	}


}