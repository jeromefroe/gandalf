package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandle(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code = %d; want 200", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error reading response: %v", err)
	}

	var (
		want   = "Hello, World!"
		actual = strings.TrimSpace(string(body))
	)
	if actual != want {
		t.Errorf("unexpected response: %s; want: %s", actual, want)
	}
}
