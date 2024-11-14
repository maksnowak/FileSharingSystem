package main

import "testing"

func TestSanity(t *testing.T) {
	if false {
		t.Fatal("expected true; got false")
	}
}

func TestFailing(t *testing.T) {
	if true {
		t.Fatal("Get lost")
	}
}
