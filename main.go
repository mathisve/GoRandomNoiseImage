package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
)

const (
	wantedSize = 50000
)

func main() {

	height := int(math.Sqrt(wantedSize)*1.15) / 2
	width := height

	min := image.Point{X: 0, Y: 0}
	max := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: min, Max: max})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, GetRandomColor())
		}
	}

	// non compressed PNG
	f, _ := os.Create("image.png")
	_ = png.Encode(f, img)

	// compressed JPEG
	f, _ = os.Create("image.jpeg")
	options := jpeg.Options{Quality: jpeg.DefaultQuality}
	_ = jpeg.Encode(f, img, &options)

	// raw bytes
	var b bytes.Buffer
	_ = png.Encode(&b, img)
	fmt.Println(b.Bytes())
}

func GetRandomColor() color.RGBA {
	return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
}
