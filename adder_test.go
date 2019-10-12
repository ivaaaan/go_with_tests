package main

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
