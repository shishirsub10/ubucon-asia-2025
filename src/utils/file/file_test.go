package utils

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFileContent_Success(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	content := "hello file"
	err := os.WriteFile(filePath, []byte(content), 0644)
	require.NoError(t, err)

	result, err := ReadFileContent(filePath)
	require.NoError(t, err)
	require.Equal(t, content, result)
}

func TestReadFileContent_FileNotFound(t *testing.T) {
	_, err := ReadFileContent("nonexistent-file.txt")
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file or directory")
}

func TestWriteFileContent_Success(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "write.txt")

	content := "hello write"
	err := WriteFileContent(filePath, content)
	require.NoError(t, err)

	// Verify file content
	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Equal(t, content, string(data))
}

func TestWriteFileContent_InvalidPath(t *testing.T) {
	invalidPath := "/invalid-dir/write.txt"
	err := WriteFileContent(invalidPath, "content")
	require.Error(t, err)
	require.Contains(t, err.Error(), "no such file or directory")
}
