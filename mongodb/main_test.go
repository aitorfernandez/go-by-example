package main

import "testing"

func TestClient(t *testing.T) {
	client := Client("mongodb://localhost:27017", "test")
	if client == nil {
		t.Errorf("client was not created %v", client)
	}
}
