package main

import (
	"fmt"

	versalog "github.com/VersaLog/VersaLog.go/VersaLog"
	Qr "github.com/kaedeek/QrGo"
)

func main() {
	logger := versalog.NewVersaLog("detailed", false, false, "QR Created", false, false, false, []string{}, false, false)
	url := "https://example.com"

	err := Qr.QrCreated(url)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to generate QR code: %v", err))
	}

	logger.Info("QR code generated successfully.")
}
