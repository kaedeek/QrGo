# QrGo

This library generates QR codes from URLs.

[japan](README.md)

## Usage
```Go
package main

import (
    "log"
    "github.com/KaedeProject/QrGo"
)

func main() {
    // URL for which you want to generate a QR code
    url := "https://example.com"

    // Generate a QR code using the library
    err := Qr.QrCreated(url)
    if err != nil {
        log.Fatalf("Failed to generate QR code: %v", err)
    }

    log.Println("QR code successfully generated.")
}

```

## Authors
- Owner and Developer: [kaedeek](https://github.com/kaedeek)
