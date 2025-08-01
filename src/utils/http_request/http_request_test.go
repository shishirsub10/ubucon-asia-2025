package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchURLContent_Success(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello from test server"))
	}))
	defer server.Close()

	content, err := FetchURLContent(server.URL)
	require.NoError(t, err)
	require.Equal(t, "hello from test server", content)
}

func TestFetchURLContent_InvalidURL(t *testing.T) {
	_, err := FetchURLContent("http://invalid-url.test.does.not-exists")
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such host")
}
