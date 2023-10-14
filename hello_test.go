package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jorge")
	want := "Hello, Jorge"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
