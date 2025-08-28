package utils

import "os/exec"

// ping executes a ping command (intentionally vulnerable due to no validation)
func Ping(ip string) (string, error) {
	cmd := exec.Command("sh", "-c", "ping -c 1 "+ip) // direct call, no sanitization
	output, err := cmd.CombinedOutput()
	return string(output), err
}
