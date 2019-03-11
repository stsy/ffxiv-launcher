package main

import (
	"fmt"
	"log"
)

func main() {

	// FIXME: Check for maintinance / worldstatus

	session, err := Login()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(session.ID)
}
