package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(6, 7)
	if got != 13 {
		t.Errorf("add(6, 7) = %d", got)
	}
}
