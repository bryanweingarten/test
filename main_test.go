package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(3, 4)
	if got != 7 {
		t.Errorf("add(3, 4) = %d", got)
	}
}
