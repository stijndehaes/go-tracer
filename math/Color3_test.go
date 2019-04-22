package math

import (
	"github.com/stretchr/testify/assert"
	gomath "math"
	"testing"
)

func TestColor3_RGBA64(t *testing.T) {
	color3 := Color3{G: 0.5, B: 1.0}
	rgba64 := color3.RGBA64()
	assert.Equal(t, uint16(0), rgba64.R)
	assert.Equal(t, uint16(gomath.MaxUint16)/2, rgba64.G)
	assert.Equal(t, uint16(gomath.MaxUint16), rgba64.B)
	assert.Equal(t, uint16(gomath.MaxUint16), rgba64.A)
}

func TestColor3_DivideFloat(t *testing.T) {
	color3 := Color3{R: 1.0, G: 1.0, B: 1.0}
	dividedColor := color3.DivideFloat(2.0)
	assert.Equal(t, 0.5, dividedColor.R)
	assert.Equal(t, 0.5, dividedColor.G)
	assert.Equal(t, 0.5, dividedColor.B)
}

func TestColor3_MultiplyFloat(t *testing.T) {
	color3 := Color3{R: 1.0, G: 1.0, B: 1.0}
	mulColor := color3.MultiplyFloat(2.0)
	assert.Equal(t, 2.0, mulColor.R)
	assert.Equal(t, 2.0, mulColor.G)
	assert.Equal(t, 2.0, mulColor.B)
}

func TestColor3_Multiply(t *testing.T) {
	color1 := Color3{R: 1.0, G: 1.0, B: 1.0}
	color2 := &Color3{R: 1.0, G: 2.0, B: 3.0}
	color3 := color1.Multiply(color2)
	assert.Equal(t, 1.0, color3.R)
	assert.Equal(t, 2.0, color3.G)
	assert.Equal(t, 3.0, color3.B)
}

func TestColor3_Add(t *testing.T) {
	color1 := Color3{R: 1.0, G: 1.0, B: 1.0}
	color2 := &Color3{R: 1.0, G: 2.0, B: 3.0}
	color3 := color1.Add(color2)
	assert.Equal(t, 2.0, color3.R)
	assert.Equal(t, 3.0, color3.G)
	assert.Equal(t, 4.0, color3.B)
}
