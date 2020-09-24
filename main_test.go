package main

import "testing"

func TestMessage(t *testing.T) {
	result := message()
	expected := "Open Source ♡ Grace Hopper!"
	if result != expected {
		t.Errorf("message() failed, expected %v, got %v", expected, result)
	}
}

func TestSquareN(t *testing.T) {
	result := squareN(2)
	expected := 4
	if result != expected {
		t.Errorf("calcSquare(1) failed, expected %v, got %v", expected, result)
	}
}
