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
	// wanted size in bytes
	// ex: 50000 is 50kb
	wantedSize = 50000
)

func main() {
	// formula to approximate height and width based on a desired size
	// only works with PNG
	height := int(math.Sqrt(wantedSize)*1.15) / 2
	width := height

	// upper left corner
	min := image.Point{X: 0, Y: 0}
	// lower right corner
	max := image.Point{X: width, Y: height}

	// creates a new images
	img := image.NewRGBA(image.Rectangle{Min: min, Max: max})

	// range over pixels and assign random color
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

// GetRandomColor returns random color.RGBA
// Alpha will always be 1
func GetRandomColor() color.RGBA {
	return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
}
