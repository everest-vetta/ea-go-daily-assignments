package main

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	rec := httptest.NewRecorder() //new recorder
	r := httptest.NewRequest("GET", "/hello", nil)

	helloHandler(rec, r)

	body := rec.Result().Body
	data, _ := io.ReadAll(body)
	assert.Equal(t, string(data), "Hello World")
}
