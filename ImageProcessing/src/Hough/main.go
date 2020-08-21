package main

import (
	"Hough/filter"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

// パラメータの定義
var inputName string = "input.png"
var outputName string = "output.png"
var thetaDiv int = 360
var rhoDiv int = 1
var thresholdLine int = 100000
var thresholdBinary uint8 = 127
var colorMask bool = false
var colorMaskR uint8 = 0
var colorMaskG uint8 = 0
var colorMaskB uint8 = 0
var thresholdColor uint8 = 20
var areaMask bool = false
var areaMaskWidth int = 300
var areaMaskHeight int = 300
var areaMaskCenterX int = 150
var areaMaskCenterY int = 150

// メイン関数
func main() {
	loadEnv()
	defer printEnv()
	img := inputImage()
	fileOut, err := os.Create(outputName)
	defer fileOut.Close()
	if err != nil {
		log.Fatal(err)
	}
	// dest := filter.DoFilter(img, filter.Binary)
	bounds := img.Bounds()
	input := filter.TypeImage{
		thetaDiv,
		rhoDiv,
		thresholdLine,
		thresholdBinary,
		colorMaskR,
		colorMaskG,
		colorMaskB,
		thresholdColor,
		areaMaskWidth,
		areaMaskHeight,
		areaMaskCenterX,
		areaMaskCenterY,
		bounds.Max,
		bounds,
		img.(*image.NRGBA),
	}
	// input.GrayScale()
	// input.Binary()
	input.ColorMask()
	dest := input.Hough()
	outputImage(dest)
}

func loadEnv() {

	err := godotenv.Load(fmt.Sprintf("%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
	inputName = os.Getenv("INPUT_NAME")
	outputName = os.Getenv("OUTPUT_NAME")
	tmp, err := strconv.ParseInt(os.Getenv("THETA_DIV"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	thetaDiv = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("RHO_DIV"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	rhoDiv = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("THRESHOLD_LINE"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	thresholdLine = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("THRESHOLD_BINARY"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	thresholdBinary = uint8(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("COLOR_MASK_R"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	colorMaskR = uint8(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("COLOR_MASK_G"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	colorMaskG = uint8(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("COLOR_MASK_B"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	colorMaskB = uint8(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("THRESHOLD_COLOR"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	thresholdColor = uint8(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("AREA_MASK_WIDTH"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	areaMaskWidth = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("AREA_MASK_HEIGHT"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	areaMaskHeight = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("AREA_MASK_CENTER_X"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	areaMaskCenterX = int(tmp)
	tmp, err = strconv.ParseInt(os.Getenv("AREA_MASK_CENTER_Y"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	areaMaskCenterY = int(tmp)
}

// PrintEnv is function
func printEnv() {
	logfile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	envlog := log.New(logfile, "[ENV]", log.LstdFlags|log.LUTC)
	envlog.Println(
		"INPUT_NAME:",
		os.Getenv("INPUT_NAME"),
		"->",
		inputName,
		"(",
		reflect.TypeOf(inputName),
		")",
	)
	envlog.Println(
		"OUTPUT_NAME:",
		os.Getenv("OUTPUT_NAME"),
		"->",
		outputName,
		"(",
		reflect.TypeOf(outputName),
		")",
	)
	envlog.Println(
		"THETA_DIV:",
		os.Getenv("THETA_DIV"),
		"->",
		thetaDiv,
		"(",
		reflect.TypeOf(thetaDiv),
		")",
	)
	envlog.Println(
		"RHO_DIV:",
		os.Getenv("RHO_DIV"),
		"->",
		rhoDiv,
		"(",
		reflect.TypeOf(rhoDiv),
		")",
	)
	envlog.Println(
		"THRESHOLD_LINE:",
		os.Getenv("THRESHOLD_LINE"),
		"->",
		thresholdLine,
		"(",
		reflect.TypeOf(thresholdLine),
		")",
	)
	envlog.Println(
		"THRESHOLD_BINARY:",
		os.Getenv("THRESHOLD_BINARY"),
		"->",
		thresholdBinary,
		"(",
		reflect.TypeOf(thresholdBinary),
		")",
	)
	envlog.Println(
		"COLOR_MASK_R:",
		os.Getenv("COLOR_MASK_R"),
		"->",
		colorMaskR,
		"(",
		reflect.TypeOf(colorMaskR),
		")",
	)
	envlog.Println(
		"COLOR_MASK_G:",
		os.Getenv("COLOR_MASK_G"),
		"->",
		colorMaskG,
		"(",
		reflect.TypeOf(colorMaskG),
		")",
	)
	envlog.Println(
		"COLOR_MASK_B:",
		os.Getenv("COLOR_MASK_B"),
		"->",
		colorMaskB,
		"(",
		reflect.TypeOf(colorMaskB),
		")",
	)
	envlog.Println(
		"THRESHOLD_COLOR:",
		os.Getenv("THRESHOLD_COLOR"),
		"->",
		thresholdColor,
		"(",
		reflect.TypeOf(thresholdColor),
		")",
	)
	envlog.Println(
		"AREA_MASK_WIDTH:",
		os.Getenv("AREA_MASK_WIDTH"),
		"->",
		areaMaskWidth,
		"(",
		reflect.TypeOf(areaMaskWidth),
		")",
	)
	envlog.Println(
		"AREA_MASK_HEIGHT:",
		os.Getenv("AREA_MASK_HEIGHT"),
		"->",
		areaMaskHeight,
		"(",
		reflect.TypeOf(areaMaskHeight),
		")",
	)
	envlog.Println(
		"AREA_MASK_CENTER_X:",
		os.Getenv("AREA_MASK_CENTER_X"),
		"->",
		areaMaskCenterX,
		"(",
		reflect.TypeOf(areaMaskCenterX),
		")",
	)
	envlog.Println(
		"AREA_MASK_CENTER_Y:",
		os.Getenv("AREA_MASK_CENTER_Y"),
		"->",
		areaMaskCenterY,
		"(",
		reflect.TypeOf(areaMaskCenterY),
		")",
	)

}

func inputImage() image.Image {
	var img image.Image
	fileIn, err := os.Open(inputName)
	defer fileIn.Close()
	if err != nil {
		log.Fatal(err)
	}
	switch filepath.Ext(inputName) {
	case ".png", ".PNG":
		img, err = png.Decode(fileIn)
		if err != nil {
			log.Fatal("error:decode1\n", err)
		} else {
			log.Println("good!\n")
		}
	case ".jpg", ".jpeg", ".JPEG":
		img, err = jpeg.Decode(fileIn)
		if err != nil {
			log.Fatal("error:decode2\n", err)
		} else {
			log.Println("good!\n")
		}
	}
	bounds := img.Bounds()
	m := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(m, m.Bounds(), img, bounds.Min, draw.Src)
	img = m
	return img
}

func outputImage(dest *image.NRGBA) {
	fileOut, err := os.Create(outputName)
	defer fileOut.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 処理結果を画像形式で出力
	switch filepath.Ext(outputName) {
	case ".png", ".PNG":
		err = png.Encode(fileOut, dest)
		if err != nil {
			log.Fatal("Failed to encode PNG gradient image.")
		} else {
			log.Println("good!\n")
		}
	case ".jpg", ".jpeg", ".JPEG":
		err = jpeg.Encode(fileOut, dest, &jpeg.Options{Quality: 100})
		if err != nil {
			log.Fatal("Failed to encode JPEG gradient image.")
		} else {
			log.Println("good!\n")
		}
	}
}
