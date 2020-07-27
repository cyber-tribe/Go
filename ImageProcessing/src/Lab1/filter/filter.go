package filter

// 必要なライブラリのインストール
import (
	"image"
	"image/color"
	"math"
	"sort"
)

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
func Binary(filter []uint32) uint32 {
	if filter[4] > 128 {
		return uint32(255)
	} else {
		return uint32(0)
	}
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
