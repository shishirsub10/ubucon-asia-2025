package utils

import "os"

// readFileContent reads the content of a given file
func ReadFileContent(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// writeFileContent writes the content to a given file path
func WriteFileContent(filePath string, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}