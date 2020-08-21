package filter

// 必要なライブラリのインストール
import (
	"image"
	"image/color"
	"math"
	"sort"
)

type Point struct{ x, y int }
type TypeImage struct {
	ThetaDiv        int
	RhoDiv          int
	ThresholdLine   int
	ThresholdBinary uint8
	ColorMaskR      uint8
	ColorMaskG      uint8
	ColorMaskB      uint8
	ThresholdColor  uint8
	AreaMaskWidth   int
	AreaMaskHeight  int
	AreaMaskCenterX int
	AreaMaskCenterY int
	Max             image.Point
	Bounds          image.Rectangle
	Img             *image.NRGBA
}
type InterfaceImage interface {
	GrayScale()
	Hough() *image.NRGBA
	Binary()
}

func DoFilter(img image.Image, fil func(filter []uint32) uint32) *image.RGBA {
	// 入力データを走査するデータの作成
	bounds := img.Bounds()
	// 出力用の入れ物を作成
	dest := image.NewRGBA(bounds)
	r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
	// 入力データを走査するデータの作成
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			filter := []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1}

			if y == 0 {
				filter[0] = 0
				filter[1] = 0
				filter[2] = 0
			}
			if x == 0 {
				filter[0] = 0
				filter[3] = 0
				filter[6] = 0
			}
			if y == bounds.Max.Y-1 {
				filter[6] = 0
				filter[7] = 0
				filter[8] = 0
			}
			if x == bounds.Max.X-1 {
				filter[2] = 0
				filter[5] = 0
				filter[8] = 0
			}
			if filter[0] != 0 {
				r, g, b, a = img.At(x-1, y-1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[0] = (r + g + b) / 3
			}

			if filter[1] != 0 {
				r, g, b, a = img.At(x, y-1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[1] = (r + g + b) / 3
			}

			if filter[2] != 0 {
				r, g, b, a = img.At(x+1, y-1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[2] = (r + g + b) / 3
			}

			if filter[3] != 0 {
				r, g, b, a = img.At(x-1, y).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[3] = (r + g + b) / 3
			}

			if filter[4] != 0 {
				r, g, b, a = img.At(x, y).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[4] = (r + g + b) / 3
			}

			if filter[5] != 0 {
				r, g, b, a = img.At(x+1, y).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[5] = (r + g + b) / 3
			}

			if filter[6] != 0 {
				r, g, b, a = img.At(x-1, y+1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[6] = (r + g + b) / 3
			}

			if filter[7] != 0 {
				r, g, b, a = img.At(x, y+1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[7] = (r + g + b) / 3
			}

			if filter[8] != 0 {
				r, g, b, a = img.At(x+1, y+1).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8
				filter[8] = (r + g + b) / 3
			}
			res := fil(filter)
			col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
			dest.Set(x, y, col)
		}
	}
	return dest
}
func DoFilter2(img image.Image, fil1 func(filter []uint32) uint32, fil2 func(filter []uint32) uint32) *image.RGBA {
	// 入力データを走査するデータの作成
	bounds := img.Bounds()
	// 出力用の入れ物を作成
	dest := image.NewRGBA(bounds)
	r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
	// 入力データを走査するデータの作成
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			filter := []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1}
			if y == 0 {
				filter[0] = 0
				filter[1] = 0
				filter[2] = 0
			}
			if x == 0 {
				filter[0] = 0
				filter[3] = 0
				filter[6] = 0
			}
			if y == bounds.Max.Y-1 {
				filter[6] = 0
				filter[7] = 0
				filter[8] = 0
			}
			if x == bounds.Max.X-1 {
				filter[2] = 0
				filter[5] = 0
				filter[8] = 0
			}
			if (y%2 == 0 && x%2 == 0) || (y%2 != 0 && x%2 != 0) {
				if filter[0] != 0 {
					r, g, b, a = img.At(x-1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[0] = (r + g + b) / 3
				}

				if filter[1] != 0 {
					r, g, b, a = img.At(x, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[1] = (r + g + b) / 3
				}

				if filter[2] != 0 {
					r, g, b, a = img.At(x+1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[2] = (r + g + b) / 3
				}

				if filter[3] != 0 {
					r, g, b, a = img.At(x-1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[3] = (r + g + b) / 3
				}

				if filter[4] != 0 {
					r, g, b, a = img.At(x, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[4] = (r + g + b) / 3
				}

				if filter[5] != 0 {
					r, g, b, a = img.At(x+1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[5] = (r + g + b) / 3
				}

				if filter[6] != 0 {
					r, g, b, a = img.At(x-1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[6] = (r + g + b) / 3
				}

				if filter[7] != 0 {
					r, g, b, a = img.At(x, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[7] = (r + g + b) / 3
				}

				if filter[8] != 0 {
					r, g, b, a = img.At(x+1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[8] = (r + g + b) / 3
				}
				res := fil1(filter)
				col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
				dest.Set(x, y, col)
			} else {
				if filter[0] != 0 {
					r, g, b, a = img.At(x-1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[0] = (r + g + b) / 3
				}

				if filter[1] != 0 {
					r, g, b, a = img.At(x, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[1] = (r + g + b) / 3
				}

				if filter[2] != 0 {
					r, g, b, a = img.At(x+1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[2] = (r + g + b) / 3
				}

				if filter[3] != 0 {
					r, g, b, a = img.At(x-1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[3] = (r + g + b) / 3
				}

				if filter[4] != 0 {
					r, g, b, a = img.At(x, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[4] = (r + g + b) / 3
				}

				if filter[5] != 0 {
					r, g, b, a = img.At(x+1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[5] = (r + g + b) / 3
				}

				if filter[6] != 0 {
					r, g, b, a = img.At(x-1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[6] = (r + g + b) / 3
				}

				if filter[7] != 0 {
					r, g, b, a = img.At(x, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[7] = (r + g + b) / 3
				}

				if filter[8] != 0 {
					r, g, b, a = img.At(x+1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[8] = (r + g + b) / 3
				}
				res := fil2(filter)
				col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
				dest.Set(x, y, col)
			}
		}
	}
	return dest
}
func DoFilter3(img image.Image, fil1 func(filter []uint32) uint32, fil2 func(filter []uint32) uint32, fil3 func(filter []uint32) uint32) *image.RGBA {
	// 入力データを走査するデータの作成
	bounds := img.Bounds()
	// 出力用の入れ物を作成
	dest := image.NewRGBA(bounds)
	r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
	// 入力データを走査するデータの作成
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			filter := []uint32{1, 1, 1, 1, 1, 1, 1, 1, 1}
			if y == 0 {
				filter[0] = 0
				filter[1] = 0
				filter[2] = 0
			}
			if x == 0 {
				filter[0] = 0
				filter[3] = 0
				filter[6] = 0
			}
			if y == bounds.Max.Y-1 {
				filter[6] = 0
				filter[7] = 0
				filter[8] = 0
			}
			if x == bounds.Max.X-1 {
				filter[2] = 0
				filter[5] = 0
				filter[8] = 0
			}
			if (y%3 == 0 && x%3 == 0) || (y%3 == 1 && x%3 == 1) || (y%3 == 2 && x%3 == 2) {
				if filter[0] != 0 {
					r, g, b, a = img.At(x-1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[0] = (r + g + b) / 3
				}

				if filter[1] != 0 {
					r, g, b, a = img.At(x, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[1] = (r + g + b) / 3
				}

				if filter[2] != 0 {
					r, g, b, a = img.At(x+1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[2] = (r + g + b) / 3
				}

				if filter[3] != 0 {
					r, g, b, a = img.At(x-1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[3] = (r + g + b) / 3
				}

				if filter[4] != 0 {
					r, g, b, a = img.At(x, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[4] = (r + g + b) / 3
				}

				if filter[5] != 0 {
					r, g, b, a = img.At(x+1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[5] = (r + g + b) / 3
				}

				if filter[6] != 0 {
					r, g, b, a = img.At(x-1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[6] = (r + g + b) / 3
				}

				if filter[7] != 0 {
					r, g, b, a = img.At(x, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[7] = (r + g + b) / 3
				}

				if filter[8] != 0 {
					r, g, b, a = img.At(x+1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[8] = (r + g + b) / 3
				}
				res := fil1(filter)
				col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
				dest.Set(x, y, col)
			} else if (y%3 == 0 && x%3 == 2) || (y%3 == 1 && x%3 == 0) || (y%3 == 2 && x%3 == 1) {
				if filter[0] != 0 {
					r, g, b, a = img.At(x-1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[0] = (r + g + b) / 3
				}

				if filter[1] != 0 {
					r, g, b, a = img.At(x, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[1] = (r + g + b) / 3
				}

				if filter[2] != 0 {
					r, g, b, a = img.At(x+1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[2] = (r + g + b) / 3
				}

				if filter[3] != 0 {
					r, g, b, a = img.At(x-1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[3] = (r + g + b) / 3
				}

				if filter[4] != 0 {
					r, g, b, a = img.At(x, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[4] = (r + g + b) / 3
				}

				if filter[5] != 0 {
					r, g, b, a = img.At(x+1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[5] = (r + g + b) / 3
				}

				if filter[6] != 0 {
					r, g, b, a = img.At(x-1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[6] = (r + g + b) / 3
				}

				if filter[7] != 0 {
					r, g, b, a = img.At(x, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[7] = (r + g + b) / 3
				}

				if filter[8] != 0 {
					r, g, b, a = img.At(x+1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[8] = (r + g + b) / 3
				}
				res := fil2(filter)
				col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
				dest.Set(x, y, col)
			} else {
				if filter[0] != 0 {
					r, g, b, a = img.At(x-1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[0] = (r + g + b) / 3
				}

				if filter[1] != 0 {
					r, g, b, a = img.At(x, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[1] = (r + g + b) / 3
				}

				if filter[2] != 0 {
					r, g, b, a = img.At(x+1, y-1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[2] = (r + g + b) / 3
				}

				if filter[3] != 0 {
					r, g, b, a = img.At(x-1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[3] = (r + g + b) / 3
				}

				if filter[4] != 0 {
					r, g, b, a = img.At(x, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[4] = (r + g + b) / 3
				}

				if filter[5] != 0 {
					r, g, b, a = img.At(x+1, y).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[5] = (r + g + b) / 3
				}

				if filter[6] != 0 {
					r, g, b, a = img.At(x-1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[6] = (r + g + b) / 3
				}

				if filter[7] != 0 {
					r, g, b, a = img.At(x, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[7] = (r + g + b) / 3
				}

				if filter[8] != 0 {
					r, g, b, a = img.At(x+1, y+1).RGBA()
					r, g, b, a = r>>8, g>>8, b>>8, a>>8
					filter[8] = (r + g + b) / 3
				}
				res := fil3(filter)
				col := color.RGBA{R: uint8(res), G: uint8(res), B: uint8(res), A: uint8(255)}
				dest.Set(x, y, col)
			}
		}
	}
	return dest
}

func (input TypeImage) Hough() *image.NRGBA {
	// 出力用の入れ物を作成
	dest := input.Img
	bounds := input.Bounds
	rhoMax := int(math.Sqrt(float64(bounds.Max.X*bounds.Max.X + bounds.Max.Y*bounds.Max.Y)))
	r, g, b, a := uint32(0), uint32(0), uint32(0), uint32(0)
	thetaMax := input.ThetaDiv
	ALPHA := input.ThresholdLine
	pins := make([][]int, thetaMax)
	for i := 0; i < thetaMax; i++ {
		pins[i] = make([]int, rhoMax)
	}
	// 入力データを走査するデータの作成
	pi := math.Pi / float64(thetaMax/2)
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a = input.Img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			col := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(255)}
			dest.Set(x, y, col)
			if r == 255 {
				continue
			}
			for i := 0; i < thetaMax; i++ {
				th := pi * float64(i)
				rho := int(float64(x)*math.Cos(th) + float64(y)*math.Sin(th))
				if 0 <= rho && rho < rhoMax {
					pins[i][rho]++
				}
			}
		}
	}
	for i := 0; i < thetaMax; i++ {
		for j := 0; j < rhoMax; j++ {
			// log.Println(pins[i][j])
			if pins[i][j] < ALPHA {
				continue
			}
			th := float64(i) * pi
			rho := float64(j)
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				if math.Sin(th) != 0 {
					y := int(-(math.Cos(th)/math.Sin(th))*float64(x) + rho/math.Sin(th))
					col := color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
					dest.Set(x, y, col)
				}
			}
			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				if math.Cos(th) != 0 {
					x := int(-(math.Sin(th)/math.Cos(th))*float64(y) + rho/math.Cos(th))
					col := color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
					dest.Set(x, y, col)
				}
			}
			for u := -60; u < 60; u++ {
				for v := -60; v < 60; v++ {
					if -1 < i+u && i+u < thetaMax && -1 < j+v && j+v < rhoMax {
						pins[i+u][j+v] = 0
					}
				}
			}
		}
	}
	// for k := 0; k < ALPHA; k++ {
	// 	max := 0
	// 	maxTh := 0
	// 	maxRho := 0
	// 	for i := 0; i < thetaMax; i++ {
	// 		for j := 0; j < rhoMax; j++ {
	// 			if pins[i][j] > max {
	// 				max = pins[i][j]
	// 				maxTh = i
	// 				maxRho = j
	// 			}
	// 		}
	// 	}
	// 	log.Println(maxTh, maxRho, pins[maxTh][maxRho])
	// 	pins[maxTh][maxRho] = 0
	// 	th := float64(maxTh) * pi
	// 	rho := float64(maxRho)
	// 	for x := bounds.Min.X; x < bounds.Max.X; x++ {
	// 		if math.Sin(th) != 0 {
	// 			y := int(-(math.Cos(th)/math.Sin(th))*float64(x) + rho/math.Sin(th))
	// 			col := color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
	// 			dest.Set(x, y, col)
	// 		}
	// 	}
	// 	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
	// 		if math.Cos(th) != 0 {
	// 			x := int(-(math.Sin(th)/math.Cos(th))*float64(y) + rho/math.Cos(th))
	// 			col := color.RGBA{R: uint8(255), G: uint8(0), B: uint8(0), A: uint8(255)}
	// 			dest.Set(x, y, col)
	// 		}
	// 	}
	// 	for i := -60; i < 60; i++ {
	// 		for j := -60; j < 60; j++ {
	// 			if -1 < maxTh+i && maxTh+i < thetaMax && -1 < maxRho+j && maxRho+j < rhoMax {
	// 				pins[maxTh+i][maxRho+j] = 0
	// 			}
	// 		}
	return dest
}
func (input TypeImage) GrayScale() {
	for x := 0; x < input.Bounds.Max.X; x++ {
		for y := 0; y < input.Bounds.Max.Y; y++ {
			r, g, b, a := input.Img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			T := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
			col := color.RGBA{R: uint8(T), G: uint8(T), B: uint8(T), A: uint8(255)}
			input.Img.Set(x, y, col)
		}
	}
}
func (input TypeImage) Binary() {
	for x := 0; x < input.Bounds.Max.X; x++ {
		for y := 0; y < input.Bounds.Max.Y; y++ {
			r, g, b, a := input.Img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			T := uint8((r + g + b) / 3)
			col := color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
			if T < input.ThresholdBinary {
				col = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
			}
			input.Img.Set(x, y, col)
		}
	}
}
func (input TypeImage) ColorMask() {
	for x := 0; x < input.Bounds.Max.X; x++ {
		for y := 0; y < input.Bounds.Max.Y; y++ {
			r, g, b, a := input.Img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			col := color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)}
			rMinus := input.ColorMaskR - input.ThresholdColor
			if int64(input.ColorMaskR)-int64(input.ThresholdColor) < 0 {
				rMinus = 0
			}
			rPlus := input.ColorMaskR + input.ThresholdColor
			if int64(input.ColorMaskR)+int64(input.ThresholdColor) > 255 {
				rPlus = 255
			}
			gMinus := input.ColorMaskG - input.ThresholdColor
			if int64(input.ColorMaskG)-int64(input.ThresholdColor) < 0 {
				gMinus = 0
			}
			gPlus := input.ColorMaskG + input.ThresholdColor
			if int64(input.ColorMaskG)+int64(input.ThresholdColor) > 255 {
				gPlus = 255
			}
			bMinus := input.ColorMaskB - input.ThresholdColor
			if int64(input.ColorMaskB)-int64(input.ThresholdColor) < 0 {
				bMinus = 0
			}
			bPlus := input.ColorMaskB + input.ThresholdColor
			if int64(input.ColorMaskB)+int64(input.ThresholdColor) > 255 {
				bPlus = 255
			}
			if rMinus <= uint8(r) && uint8(r) <= rPlus && gMinus <= uint8(g) && uint8(g) <= gPlus && bMinus <= uint8(b) && uint8(b) <= bPlus {
				col = color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(255)}
			}
			input.Img.Set(x, y, col)
		}
	}
}
func (input TypeImage) Prewitt_V() {
}
func None(filter []uint32) uint32 {
	return filter[4]
}
func Prewitt_V(filter []uint32) uint32 {
	res := float64(0)
	res += -float64(filter[0])
	res += 0
	res += float64(filter[2])
	res += -float64(filter[3])
	res += 0
	res += float64(filter[5])
	res += -float64(filter[6])
	res += 0
	res += float64(filter[8])

	return uint32(math.Abs(res))
}
func Prewitt_H(filter []uint32) uint32 {
	res := float64(0)
	res += -float64(filter[0])
	res += -float64(filter[1])
	res += -float64(filter[2])
	res += 0
	res += 0
	res += 0
	res += float64(filter[6])
	res += float64(filter[7])
	res += float64(filter[8])

	return uint32(math.Abs(res))
}

func Average(filter []uint32) uint32 {
	res := float64(0)
	res += float64(filter[0])
	res += float64(filter[1])
	res += float64(filter[2])
	res += float64(filter[3])
	res += float64(filter[4])
	res += float64(filter[5])
	res += float64(filter[6])
	res += float64(filter[7])
	res += float64(filter[8])
	res /= 9

	return uint32(math.Abs(res))
}
func Median(filter []uint32) uint32 {
	tmp := filter[:]
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})

	mNumber := len(tmp) / 2

	if mNumber%2 != 0 {
		return uint32(tmp[mNumber])
	}

	return uint32((tmp[mNumber-1] + tmp[mNumber]) / 2)
}
func Nega(filter []uint32) uint32 {
	return uint32(math.Abs(float64(255) - float64(filter[4])))
}
