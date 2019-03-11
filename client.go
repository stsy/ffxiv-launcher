package main

import (
	"fmt"
	"log"
)

// Login checks if logon info is stored in config
// If not: Promt
func Login() {
	config, err := config.Load("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	// FIXME: Check if session is stored

	var token string
	if config.Auth.UserID == "" {
		fmt.Print("Username: ")
		fmt.Scanln(&config.Auth.UserID)
	}

	if config.Auth.Password == "" {
		fmt.Print("Password: ")
		fmt.Scanln(&config.Auth.Password)
	}

	if config.Auth.Token {
		fmt.Print("Token: ")
		fmt.Scanln(&token)
	}

	// FIXME: Autologin? Return?
}
