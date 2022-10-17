package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
