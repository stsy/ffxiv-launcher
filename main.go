package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("FFXIV Launcher v0.3")
	fmt.Println("https://github.com/stsy/ffxiv-launcher")
	fmt.Println("")

	// Print what expansion

	// Check world status
	if err := worldStatus(); err != nil {
		fmt.Println("ERROR")
		fmt.Println("")
	}

	// Attempt to autologin if enabled
	if autoLogin() {
		// Sleep for 8 sec before closing program
		duration := time.Duration(8) * time.Second
		time.Sleep(duration)
		return
	}

	fmt.Println("")

	// Login loop
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
