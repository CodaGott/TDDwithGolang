package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if want != got{
		t.Errorf("got %q want %q", got, want)
	}
}
