package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Huds")
	want := "Hello, Huds"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
