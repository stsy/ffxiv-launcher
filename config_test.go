package main

import "testing"

// Test config.Load()
func TestConfig(t *testing.T) {
	config, err := config.Load("./config/config.json")
	got := config.Hello
	expected := "world"

	if got != expected {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
	if err != nil {
		t.Error(err)
	}
}
