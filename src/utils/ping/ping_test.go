package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing_InvalidIP(t *testing.T) {
	_, err := Ping("256.256.256.256") // Invalid IP
	require.Error(t, err)
}

func TestPing_Localhost(t *testing.T) {
	output, err := Ping("127.0.0.1")
	require.NoError(t, err)
	require.Contains(t, output, "PING")
}
