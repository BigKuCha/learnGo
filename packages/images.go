package packages

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	width  = 510
	height = 510
)

func NewPngImg() {
	file, _ := os.Create("test.png")
	defer file.Close()
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Println(i % 255)
			rgba := color.RGBA{uint8(i % 255), uint8(j % 255), uint8((i * j) % 255), 255}
			img.SetRGBA(i, j, rgba)
		}
	}
	png.Encode(file, img)
}
