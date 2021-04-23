package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Denis")
	want := "Hello, Denis"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
