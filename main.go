package main

import (
	"fmt"
	"log"
)

func main() {

	// TODO: Setup

	fmt.Println("FFXIV Launcher v0.2")
	fmt.Println("https://github.com/stsy/ffxiv-launcher")
	fmt.Println("")

	// Check world status
	if err := worldStatus(); err != nil {
		fmt.Println("ERROR")
		fmt.Println("")
	}

	for {
		// Login or retry
		session, err := Login()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Launch game
		if err := Launcher(session); err != nil {
			log.Fatal(err)
		}

		fmt.Println("Starting game ..")
		break
	}
}
