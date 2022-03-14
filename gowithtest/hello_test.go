package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to a specified name", func(t *testing.T) {
		got := Hello("Dozie", "Spanish")
		want := "Hola, Dozie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in french", func(t *testing.T) {
		got := Hello("Emma", "French")
		want := "Bonjour, Emma"
		assertCorrectMessage(t, got, want)
	})
}

func ExampleHello() {
	hello := Hello("Dozie", "English")
	fmt.Println(hello)
	// Output: Hello, Dozie
}
