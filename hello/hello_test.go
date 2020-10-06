package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kyle")
	want := "Hello, Kyle!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
