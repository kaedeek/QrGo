# QrGo

このライブラリは、urlをQRコードにするライブラリです。

[English](README-en.md)

## Usage
```Go
package main

import (
    "log"
    "github.com/KaedeProject/QrGo"
)

func main() {
	// QRコードを生成したいURL
	url := "https://example.com"

	// ライブラリを使ってQRコードを生成
	err := Qr.QrCreated(url)
	if err != nil {
		log.Fatalf("QRコード生成に失敗しました: %v", err)
	}

	log.Println("QRコードが正常に生成されました。")
}
```

## Authors
- Owner and Developer: [kaedeek](https://github.com/kaedeek)
