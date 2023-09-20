package main

import (
	"fmt"
	"log"

	"github.com/bassman7689/fetch-exercise/pkg/config"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("err: %+v\n", err)
	}
	fmt.Printf("%+v\n", conf)
}
