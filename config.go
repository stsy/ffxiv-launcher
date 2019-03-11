package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Init
var config *Config

// Config struct
type Config struct {
	Hello string `json:"hello"`
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
