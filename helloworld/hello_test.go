package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q",got, want)
		}	
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gastón", "")
		want := "Hello, Gastón"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		
		assertCorrectMessage(t, got, want)
	})

	t.Run("Given Spanish language say hello in that language", func(t *testing.T) {
		got := Hello("Gastón", "Spanish")
		want := "Hola, Gastón"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Given another language say hello in that oher language", func(t *testing.T) {
		got := Hello("Gastón", "French")
		want := "Bonjour, Gastón"

		assertCorrectMessage(t, got, want)
	})

}
