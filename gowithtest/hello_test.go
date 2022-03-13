package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Dozie")
	want := "Hello, Dozie"

	if want != got{
		t.Errorf("got %q want %q", got, want)
	}
}
