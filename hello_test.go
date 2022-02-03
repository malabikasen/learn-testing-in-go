package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, expected string, returned string) {
		t.Helper()
		if expected != returned {
			t.Errorf("expected = %q, return value = %q", expected, returned)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		returnedValue := PrintHello("Chris", "")
		expectedValue := "Hello, Chris"
		assertCorrectMessage(t, expectedValue, returnedValue)

	})

	t.Run("saying hello world when empty string is supplied", func(t *testing.T) {
		returnedValue := PrintHello("", "")
		expectedValue := "Hello, World"
		assertCorrectMessage(t, expectedValue, returnedValue)
	})

	t.Run("spanish greeting", func(t *testing.T) {
		returnedValue := PrintHello("Senorita", "Spanish")
		expectedValue := "Hola, Senorita"
		assertCorrectMessage(t, expectedValue, returnedValue)
	})

	t.Run("french greeting", func(t *testing.T) {
		returnedValue := PrintHello("Emily", "French")
		expectedValue := "Bonjour, Emily"
		assertCorrectMessage(t, expectedValue, returnedValue)
	})

}
