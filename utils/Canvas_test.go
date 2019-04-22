package utils

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stretchr/testify/assert"
	"image"
	gomath "math"
	"testing"
)

func TestPut(t *testing.T) {
	canvas := CreateCanvas(image.Point{X: 100, Y: 100})
	color := math.Color3{R: 1.0, G: 0.5, B: 0.0}
	canvas.Put(0, 100, &color)
	assert.Equal(t, uint16(gomath.MaxUint16), canvas.image.RGBA64At(0, 0).R)
	assert.Equal(t, uint16(gomath.MaxUint16/2), canvas.image.RGBA64At(0, 0).G)
	assert.Equal(t, uint16(0), canvas.image.RGBA64At(0, 100).B)
}
