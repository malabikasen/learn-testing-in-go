package main

import "testing"

func TestHello(t *testing.T) {
	returnedValue := PrintHello("Chris")
	expectedValue := "Hello, Chris"

	if returnedValue != expectedValue {
		t.Errorf("Expected %q, got %q", expectedValue, returnedValue)
	}
}
