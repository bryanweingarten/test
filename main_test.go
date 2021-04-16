package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(4, 5)
	if got != 9 {
		t.Errorf("add(4, 5) = %d", got)
	}
}
