package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandleExample(t *testing.T) {
	want := "this is my response\n"

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handleExample(w, req)

	res := w.Result()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := string(body)
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
