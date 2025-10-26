package Qr

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"

	versalog "github.com/VersaLog/VersaLog.go/VersaLog"
)

func QrCreated(url string) error {
	logger := versalog.NewVersaLog("detailed", false, true, "QR Created", false, false, false, []string{}, false, false)
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
		logger.Error(fmt.Sprintf("%s\ninfo: %s\n", err, stderr.String()))
		return err
	}

	logger.Info("successfully: %s\n", out.String())
	return nil
}
