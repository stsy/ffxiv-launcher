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
	client, err := client.Load(clientPath)
	if err != nil {
		return
	}

	// GET request to oauth /top
	req, err := http.NewRequest("GET", client.Launcher.Oauth.Get, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", client.Launcher.UserAgent)

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
	m, err := ReSearch(client.Launcher.Oauth.RegexStored,
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
	user, err := user.Load(userPath)
	if err != nil {
		return
	}

	client, err := client.Load(clientPath)
	if err != nil {
		return
	}

	// FIXME: Check if session is stored

	var token string
	if user.Auth.UserID != "" {
		fmt.Println("Username: " + user.Auth.UserID)
	} else {
		fmt.Print("Username: ")
		fmt.Scanln(&user.Auth.UserID)
	}

	if user.Auth.Password == "" {
		fmt.Print("Password: ")
		fmt.Scanln(&user.Auth.Password)
	}

	if user.Auth.Token {
		fmt.Print("Token: ")
		fmt.Scanln(&token)
	}

	stored, err := Stored()
	if err != nil {
		return
	}

	payload := fmt.Sprintf(`_STORED_=%s&sqexid=%s&password=%s&otppw=%s`,
		stored, user.Auth.UserID, user.Auth.Password, token)

	// POST stored, user, pass, token to /login.send
	body := strings.NewReader(payload)
	req, err := http.NewRequest("POST", client.Launcher.Oauth.Post, body)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", client.Launcher.UserAgent)
	req.Header.Set("Referer", client.Launcher.Oauth.Post)
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

	m, err := ReSearch(client.Launcher.Oauth.RegexSid, string(source))
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
	client, err := client.Load(clientPath)
	if err != nil {
		return
	}

	// Make filehash payload
	var payload string
	for i, file := range client.Game.Files {
		hash, err := Hash(client.Game.Path.Boot + file)
		if err != nil {
			log.Fatal(err)
		}
		payload += fmt.Sprintf("%s/%s", file, hash)
		if i < len(client.Game.Files)-1 {
			payload += ","
		}
	}

	// Read out current game version form ffxivgame.ver file
	version, err := ioutil.ReadFile(client.Game.Path.Game + "ffxivgame.ver")
	if err != nil {
		return
	}

	// Disable SSL check
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	hc := &http.Client{Transport: tr}

	// Make gamever url from version and sid
	url := fmt.Sprintf(client.Game.GameverURL, string(version), s.ID)

	// Send payload
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Hash-Check", "enabled")

	resp, err := hc.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Launch args
	// FIXME: add some of this to config
	args := []string{
		"DEV.TestSID=" + resp.Header.Get("X-Patch-Unique-Id"),
		"DEV.UseSqPack=1",
		"DEV.DataPathType=1",
		"DEV.MaxEntitledExpansionID=" + client.Game.Expansion,
		"SYS.Region=3",
		"language=1",
	}

	// Check if DX11 is enabled
	bin := "ffxiv.exe"
	if client.Game.Dx11 {
		bin = "ffxiv_dx11.exe"
	}

	// Start game
	Start(client.Game.Path.Game+bin, args)

	return
}
