package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config *Config

// Config struct
type Config struct {
	Auth struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
		Token    bool   `json:"token"`
		Session  struct {
			Date string `json:"date"`
			ID   string `json:"id"`
		} `json:"session"`
	} `json:"auth"`
	Launcher struct {
		UserAgent string `json:"user_agent"`
		Oauth     struct {
			Regex string `json:"regex"`
			Get   string `json:"get"`
			Post  string `json:"post"`
		} `json:"oauth"`
	} `json:"launcher"`
	Testing string `json:"testing"`
}

// Load config from filepath
// eg. "./config/config.json"
func (*Config) Load(path string) (c *Config, err error) {
	// Check if config file exists, if not create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// FIXME: create config.json if it doesn't exitst
	}

	j, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	json.Unmarshal(j, &c)
	return
}
