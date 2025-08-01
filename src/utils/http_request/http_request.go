package utils

import (
	"io"
	"net/http"
)

// fetchURLContent sends a GET request to the given URL and returns the response body
func FetchURLContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
