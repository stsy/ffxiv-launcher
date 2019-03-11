package main

import (
	"fmt"
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
