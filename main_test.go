package main

import "testing"

func TestHello(t *testing.T) {
	emptyResult := hello("")
	expected := "Open Source ♡ Grace Hopper!"
	if emptyResult != expected {
		t.Errorf("hello(\"\") failed, expected %v, got %v", expected, emptyResult)
	}
}
