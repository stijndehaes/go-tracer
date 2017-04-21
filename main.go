package main

import (
	"github.com/Steniaz/go-tracer/utils"
	"image"
	"github.com/Steniaz/go-tracer/math"
	"github.com/Steniaz/go-tracer/raytracer"
	"github.com/Steniaz/go-tracer/geometry"
	"github.com/Steniaz/go-tracer/accelerators"
)

func main() {
	canvas := utils.CreateCanvas(image.Point{100, 100})
	camera := raytracer.GetCamera(
		math.Vector3{X: -10},
		math.Vector3{Y: 1.0},
		math.Vector3{X: 1.0},
		90.0,
		1.0,
		100,
		100,
	)
	sphere :=  geometry.Sphere{
		math.Vector3{},
		1.0,
	}
	acc := accelerators.DumbAccelerator{}

	acc.AddGeometry(&sphere)


	for x := 0; x < canvas.Size.X; x++ {
		for y := 0; y < canvas.Size.Y; y++ {
			ray := camera.GenerateRay(x, y)
			acc.IntersectGeometry(&ray)
			if ray.Hit {
				canvas.Put(x, y, math.Color3{ R: 1.0 })
			}

		}
	}
	canvas.SaveAsFile("out.png")
}