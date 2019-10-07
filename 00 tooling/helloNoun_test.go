package main

import "testing"

func TestHelloNoun(t *testing.T) {
	testMsg := "hello foobar"
	result := helloNoun("foobar")

	if result != testMsg {
		t.Fatalf("%s != %s", result, testMsg)
	}
}
