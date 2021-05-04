package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(5, 5)
	if got != 10 {
		t.Errorf("add(6, 7) = %d", got)
	}
}
