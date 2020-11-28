package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAddGet(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/add/1/1")

	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestProductGet(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/product/1/1")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)

}

type testNumbers struct {
	number1 int
	number2 int
}

func TestAdd(t *testing.T) {

	numbers := []testNumbers{
		testNumbers{1, 1},
		testNumbers{2, 1000},
		testNumbers{333, 333},
	}
	want := []int{2, 1002, 666}
	for i, test := range numbers {
		result := add(test.number1, test.number2)
		if result != want[i] {
			t.Errorf("Got = %v; want %v", result, want[i])
		}
	}
}

func TestMultiply(t *testing.T) {

	numbers := []testNumbers{
		testNumbers{1, 1},
		testNumbers{2, 1000},
		testNumbers{5, 333},
	}
	want := []int{1, 2000, 1665}
	for i, test := range numbers {
		result := multiply(test.number1, test.number2)
		if result != want[i] {
			t.Errorf("Got = %v; want %v", result, want[i])
		}
	}
}
