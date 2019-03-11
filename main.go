package main

import (
	"fmt"
	"log"
)

func main() {
	config, err := config.Load("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.Hello)
}
