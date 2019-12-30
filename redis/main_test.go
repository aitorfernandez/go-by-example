package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	Set("mykey", "hello")

	mykey := Get("mykey")
	if mykey != "hello" {
		t.Errorf("mykey was incorrect %v", mykey)
	}
}

func TestHget(t *testing.T) {
	Hset("myhash", "field", "hello")

	field := Hget("myhash", "field")
	if field != "hello" {
		t.Errorf("field was incorrect %v", field)
	}
}
