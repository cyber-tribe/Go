package main

// 必要なライブラリのインストール
import (
	"image"
	"image/color"
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

	// 入力データを走査するデータの作成
	bounds := img.Bounds()
	// 出力用の入れ物を作成
	dest := image.NewRGBA(bounds)
	// 画像を1ピクセルずつ走査
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// ピクセルの取得
			curPixel := img.At(x, y)
			// ピクセルから色情報RGBA(16bit)を取得
			r, g, b, a := curPixel.RGBA()
			// 色情報16bitを8bitの形式に変換
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			/*　基本的なべイヤー配列

			緑　赤　緑　赤　緑　赤

			青　緑　青　緑　青　緑

			緑　赤　緑　赤　緑　赤

			青　緑　青　緑　青　緑

			緑　赤　緑　赤　緑　赤

			青　緑　青　緑　青　緑

			*/
			if y%2 == 0 {
				if x%2 == 0 {
					// 緑
					col := color.RGBA{R: uint8(0), G: uint8(g), B: uint8(0), A: uint8(255)}
					dest.Set(x, y, col)
				} else {
					// 赤
					col := color.RGBA{R: uint8(r), G: uint8(0), B: uint8(0), A: uint8(255)}
					dest.Set(x, y, col)
				}
			} else {
				if x%2 == 0 {
					// 青
					col := color.RGBA{R: uint8(0), G: uint8(0), B: uint8(b), A: uint8(255)}
					dest.Set(x, y, col)
				} else {
					// 緑
					col := color.RGBA{R: uint8(0), G: uint8(g), B: uint8(0), A: uint8(255)}
					dest.Set(x, y, col)
				}
			}
		}
	}

	// 処理結果を画像形式で出力
	err = png.Encode(os.Stdout, dest)
	if err != nil {
		panic("Failed to encode JPEG gradient image.")
	}
}
