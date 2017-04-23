package utils

import (
	"image"
	"github.com/Steniaz/go-tracer/math"
	"os"
	"image/png"
)

type Canvas struct {
	image *image.RGBA64
	Size  image.Point
}

func CreateCanvas(size image.Point) *Canvas {
	im := image.NewRGBA64(image.Rectangle{image.Point{0, 0}, size})
	return &Canvas{
		im,
		size,
	}
}

func (canvas *Canvas) Put(x, y int, color *math.Color3) {
	canvas.image.Set(x, canvas.Size.Y - y, color.RGBA64())
}

func (canvas *Canvas) SaveAsFile(filename string) {
	toimg, _ := os.Create(filename)
	png.Encode(toimg, canvas.image)
	toimg.Close()
}