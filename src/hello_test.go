package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("hj")
		want := "hello hj"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string default None", func(t *testing.T) {
		got := Hello("")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})
}