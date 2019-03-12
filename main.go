package main

import (
	"log"
)

func main() {

	// FIXME: Setup
	// FIXME: Check for maintinance / worldstatus
	// FIXME: Retry if wrong password

	// Login
	session, err := Login()
	if err != nil {
		log.Fatal(err)
	}

	// Start game
	if Launcher(session) != nil {
		log.Fatal(err)
	}
}
