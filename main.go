package main

import (
	"log"
)

func main() {

	// FIXME: Check for maintinance / worldstatus

	session, err := Login()
	if err != nil {
		log.Fatal(err)
	}

	// Start game?
	if Launcher(session) != nil {
		log.Fatal(err)
	}
}
