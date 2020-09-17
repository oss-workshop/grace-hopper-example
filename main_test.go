package main

import "testing"

func TestMessage(t *testing.T) {
	result := message()
	expected := "Open Source ♡ Grace Hopper!"
	if result != expected {
		t.Errorf("message() failed, expected %v, got %v", expected, result)
	}
}
