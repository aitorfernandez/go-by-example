package main_test

import (
	"net/http"
	"testing"
)

func TestGetPost(t *testing.T) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080/api/v1/posts/1", nil)

	res, _ := client.Do(req)

	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d want %d", res.StatusCode, http.StatusOK)
	}
}
