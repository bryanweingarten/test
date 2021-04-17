package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(5, 6)
	if got != 11 {
		t.Errorf("add(5, 6) = %d", got)
	}
}
