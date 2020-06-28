package main

import "testing"

func TestHello(t *testing.T) {
	//define expect & actual
	got := Hello("Ken")
	want := "Hello, Ken"
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}
