package math

import (
	gocolor "image/color"
	gomath "math"
)
type Color3 struct {
	R, G, B float64
}

func (color Color3) RGBA64() gocolor.RGBA64 {
	return gocolor.RGBA64{
		R: uint16(gomath.Min(color.R* gomath.MaxUint16, gomath.MaxUint16)),
		G: uint16(gomath.Min(color.G* gomath.MaxUint16, gomath.MaxUint16)),
		B: uint16(gomath.Min(color.B* gomath.MaxUint16, gomath.MaxUint16)),
		A: gomath.MaxUint16,
	}
}