package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
}

// Stored returns _STORED_ ID
func Stored() (stored string, err error) {
	config, err := config.Load("./config/config.json")
	if err != nil {
		return
	}

	// GET request to oauth GET_URL
	req, err := http.NewRequest("GET", config.Launcher.Oauth.Get, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", config.Launcher.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read response
	source, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the ID in source
	m, err := ReSearch(config.Launcher.Oauth.Regex,
		string(source))
	if err != nil {
		return
	}

	stored = m["_STORED_"]
	return

}
