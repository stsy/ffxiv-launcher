package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Session ..
type Session struct {
	ID string
}

// Stored returns _STORED_ ID
func Stored() (stored string, err error) {
	config, err := config.Load(configPath)
	if err != nil {
		return
	}

	// GET request to oauth /top
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
		return
	}

	// Find the ID in source
	m, err := ReSearch(config.Launcher.Oauth.RegexStored,
		string(source))
	if err != nil {
		return
	}

	stored = m["_STORED_"]
	return
}

// Login checks if logon info is stored in config
// If not: Promt
func Login() (s *Session, err error) {
	config, err := config.Load(configPath)
	if err != nil {
		return
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

	stored, err := Stored()
	if err != nil {
		return
	}

	payload := fmt.Sprintf(`_STORED_=%s&sqexid=%s&password=%s&otppw=%s`,
		stored, config.Auth.UserID, config.Auth.Password, token)

	// POST stored, user, pass, token to /login.send
	body := strings.NewReader(payload)
	req, err := http.NewRequest("POST", config.Launcher.Oauth.Post, body)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", config.Launcher.UserAgent)
	req.Header.Set("Referer", config.Launcher.Oauth.Post)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Read source
	source, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	m, err := ReSearch(config.Launcher.Oauth.RegexSid, string(source))
	if err != nil {
		err = fmt.Errorf("Wrong username or password")
		return
	}

	s = &Session{
		ID: m["SID"],
	}

	return
}

// Launcher ..
func Launcher(s *Session) (err error) {
	config, err := config.Load(configPath)
	if err != nil {
		return
	}

	// Make filehash payload
	var payload string
	for i, file := range config.Game.Files {
		hash, err := Hash(config.Game.Path.Boot + file)
		if err != nil {
			log.Fatal(err)
		}
		payload += fmt.Sprintf("%s/%s", file, hash)
		if i < len(config.Game.Files)-1 {
			payload += ","
		}
	}

	// Read out current game version form ffxivgame.ver file
	version, err := ioutil.ReadFile(config.Game.Path.Game + "ffxivgame.ver")
	if err != nil {
		return
	}

	// Disable SSL check
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Make gamever url from version and sid
	url := fmt.Sprintf(config.Game.GameverURL, string(version), s.ID)

	// Send payload
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Hash-Check", "enabled")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// FIXME
	fmt.Println(resp.Header.Get("X-Patch-Unique-Id"))
	return
}
