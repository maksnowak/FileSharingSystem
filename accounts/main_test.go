package main

import "testing"

func TestPassing(t *testing.T) {
	if false {
		t.Fatal("1 does indeed equal 2")
	}
}
