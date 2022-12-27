package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

const VERSION = "VERSION"

func TestUpdateHeader(t *testing.T) {
	mockRequestHeader := http.Header{}
	mockRequestHeader.Add("request", "1")

	mockResponseHeader := http.Header{}
	mockResponseHeader.Add("response", "2")

	UpdateHeader(mockResponseHeader, mockRequestHeader)
	assert.Equal(t, "1", mockResponseHeader.Get("request"))
	assert.Equal(t, "2", mockResponseHeader.Get("response"))
}

func TestGetVersion(t *testing.T) {
	origin := os.Getenv(VERSION)

	if len(origin) == 0 {
		assert.Equal(t, "default", GetVersion("default"))

		os.Setenv(VERSION, "mock")

		assert.Equal(t, "mock", GetVersion("default"))

		os.Unsetenv(VERSION)
	} else {
		assert.Equal(t, origin, GetVersion("default"))
	}
}
