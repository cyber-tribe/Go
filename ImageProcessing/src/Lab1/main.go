package main

import (
	"Lab1/filter"
	"image/png"
	"log"
	"os"
)

// メイン関数
func main() {
	// threshold := 1.5
	// 入力データの読み込み
	img, err := png.Decode(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	dest := filter.DoFilter3(img, filter.Prewitt_V, filter.Prewitt_H, filter.Nega)
	// dest := filter.DoFilter(img, filter.Nega)
	// 処理結果を画像形式で出力
	err = png.Encode(os.Stdout, dest)
	if err != nil {
		panic("Failed to encode JPEG gradient image.")
	}
}
