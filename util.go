package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

// ReSearch FindStringSubmatch, returns match map
// Regex example: `name="_STORED_" value="(?P<PARAMNAME>.*)"`
// A successful call returns err == nil
func ReSearch(pattern, input string) (map[string]string, error) {
	var compRegEx = regexp.MustCompile(pattern)
	match := compRegEx.FindStringSubmatch(input)
	if len(match) == 0 {
		return nil, fmt.Errorf("ReSearch: No match found\nPattern: %s", pattern)
	}

	params := make(map[string]string)
	for i, key := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[key] = match[i]
		}
	}
	return params, nil
}

// Hash returns "<filesize>/<sha1hash>" from given path
// A successful call returns err == nil
func Hash(path string) (sizeHash string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	fs, err := file.Stat()
	if err != nil {
		return
	}
	defer file.Close()

	hash := sha1.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}

	b := hash.Sum(nil)
	sizeHash = fmt.Sprintf("%d/%s", fs.Size(), hex.EncodeToString(b))
	return
}

// Start starts a program with args and detach
func Start(path string, args []string) {
	cmd := exec.Command(path, args...)
	cmd.Start()
}

// DownloadString returns the sourcecode form URL.
// A successful call returns err == nil.
func DownloadString(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	html = string(bytes)
	return
}
