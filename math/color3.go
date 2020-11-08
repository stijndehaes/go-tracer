package math

import (
	gocolor "image/color"
	gomath "math"
)

type Color3 struct {
	R, G, B float64
}

func (color *Color3) RGBA64() gocolor.RGBA64 {
	return gocolor.RGBA64{
		R: uint16(gomath.Max(0, gomath.Min(color.R*gomath.MaxUint16, gomath.MaxUint16))),
		G: uint16(gomath.Max(0, gomath.Min(color.G*gomath.MaxUint16, gomath.MaxUint16))),
		B: uint16(gomath.Max(0, gomath.Min(color.B*gomath.MaxUint16, gomath.MaxUint16))),
		A: gomath.MaxUint16,
	}
}

func (color *Color3) DivideFloat(float float64) *Color3 {
	return color.MultiplyFloat(1 / float)
}

func (color *Color3) MultiplyFloat(float float64) *Color3 {
	return &Color3{R: color.R * float, G: color.G * float, B: color.B * float}
}

func (color *Color3) Multiply(color2 *Color3) *Color3 {
	return &Color3{R: color.R * color2.R, G: color.G * color2.G, B: color.B * color2.B}
}

func (color *Color3) Add(color2 *Color3) *Color3 {
	return &Color3{R: color.R + color2.R, G: color.G + color2.G, B: color.B + color2.B}
}
