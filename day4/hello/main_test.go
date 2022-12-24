package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func testHelloHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/hello", nil)
	helloHandler(rec, req)

	body := rec.Result().Body
	data, _ := ioutil.ReadAll(body)

	assert.Equal(t, string(data), "Hello")
}
