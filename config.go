package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var (
	client     *Client
	user       *User
	clientPath = "./config/client.json"
	userPath   = "./config/user.json"
)

// User settings
type User struct {
	Auth struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
		Token    bool   `json:"token"`
		Session  struct {
			Date string `json:"date"`
			ID   string `json:"id"`
		} `json:"session"`
	} `json:"login"`
}

// Client settings
type Client struct {
	Launcher struct {
		UserAgent string `json:"user_agent"`
		Oauth     struct {
			RegexStored string `json:"regex_stored"`
			RegexSid    string `json:"regex_sid"`
			Get         string `json:"get"`
			Post        string `json:"post"`
		} `json:"oauth"`
	} `json:"launcher"`
	Game struct {
		Expansion string `json:"expansion"`
		Dx11      bool   `json:"dx11"`
		Path      struct {
			Boot string `json:"boot"`
			Game string `json:"game"`
		} `json:"path"`
		Files      []string `json:"files"`
		GameverURL string   `json:"gamever_url"`
	} `json:"game"`
}

// Load client config from filepath
// eg. "./config/config.json"
func (*Client) Load(path string) (c *Client, err error) {
	// Check if config file exists, if not create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// FIXME: create client.json if it doesn't exitst
	}

	j, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	json.Unmarshal(j, &c)
	return
}

// Load user config from filepath
// eg. "./config/config.json"
func (*User) Load(path string) (c *User, err error) {
	// Check if config file exists, if not create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// FIXME: create user.json if it doesn't exitst
	}

	j, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	json.Unmarshal(j, &c)
	return
}
