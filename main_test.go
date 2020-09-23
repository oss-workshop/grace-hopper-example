package main

import "testing"

func TestMessage(t *testing.T) {
	result := message()
	expected := "Open Source ♡ Grace Hopper!"
	if result != expected {
		t.Errorf("message() failed, expected %v, got %v", expected, result)
	}
}

func TestCalcSquare(t *testing.T) {
	result := calcSquare(1)
	expected := 1
	if result != expected {
		t.Errorf("calcSquare(1) failed, expected %v, got %v", expected, result)
	}
}
