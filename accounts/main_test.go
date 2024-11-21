package main

import "testing"

func TestPassing(t *testing.T) {
	if false {
		t.Fatal("1 does indeed equal 2")
	}
}

func TestAddOne(t *testing.T) {
	x := 7
	if addOne(&x) != 8 {
		t.Fatal("7 + 1 does not equal 8")
	}
}
