package main

import (
	"fmt"
	"log"
)

func main() {

	// TODO: Setup
	// TODO: Check for maintinance / worldstatus
	// https://frontier.ffxiv.com/worldStatus/current_status.json
	// https://frontier.ffxiv.com/worldStatus/gate_status.json

	fmt.Println("FFXIV Launcher")
	fmt.Println("https://github.com/stsy/ffxiv-launcher")
	fmt.Println("")

	// Login
	session, err := Login()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Retry if wrong password

	// Start game
	if Launcher(session) != nil {
		log.Fatal(err)
	}
}
