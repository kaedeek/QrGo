package Qr

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func QrCreated(url string) error {
	scriptPath, err := filepath.Abs("./base/body.py")
	if err != nil {
		return fmt.Errorf("Failed to retrieve script path: %w", err)
	}

	cmd := exec.Command("python3", scriptPath, url)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		fmt.Printf("Error: %s\ninfo: %s\n", err, stderr.String())
		return err
	}

	fmt.Printf("successfully: %s\n", out.String())
	return nil
}
