package main

import "testing"

func TestStored(t *testing.T) {
	stored, err := Stored()
	if err != nil {
		t.Error(err)
	}
	got := len(stored)
	expected := 546

	if got != expected {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
}
