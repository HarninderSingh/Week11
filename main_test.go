package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 5)
	expected := 7
	if result != expected {
		t.Fatalf("Add(2,5) = %d expected = %d", result, expected)

	}
}
func TestSubtract(t *testing.T) {
	result := Subtract(2, 5)
	expected := -3
	if result != expected {
		t.Fatalf("Subtract(2,5) = %d expected = %d", result, expected)

	}
}
func TestMultiply(t *testing.T) {
	result := Multiply(2, 5)
	expected := 10
	if result != expected {
		t.Fatalf("Multiply(2,5) = %d expected = %d", result, expected)

	}
}
