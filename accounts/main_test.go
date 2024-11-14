package main

import "testing"

func TestPassing(t *testing.T) {
	if 1 == 2 {
		t.Fatal("1 does indeed equal 2")
	}
}
