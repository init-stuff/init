package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func RunPythonCode(code string) (string, error) {
	cmd := exec.Command("py", "-c", code)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v\n%s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
