package main

import (
	"log"

	"github.com/bassman7689/fetch-exercise/pkg/config"
	"github.com/bassman7689/fetch-exercise/pkg/server"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}

	srv, err := server.New(conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}
}
