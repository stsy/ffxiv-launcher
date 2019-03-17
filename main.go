package main

import (
	"fmt"
	"log"
)

func main() {

	// TODO: Setup

	fmt.Println("FFXIV Launcher v0.1")
	fmt.Println("https://github.com/stsy/ffxiv-launcher")
	fmt.Println("")

	// Check world status
	if err := worldStatus(); err != nil {
		log.Panic(err)
	}

	// Login
	session, err := Login()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Retry if wrong password

	// Start game
	if err = Launcher(session); err != nil {
		log.Fatal(err)
	}
}
