package scenes

import (
	"github.com/stijndehaes/go-tracer/geometry"
	"github.com/stijndehaes/go-tracer/lights"
	"github.com/stijndehaes/go-tracer/materials"
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
)

var s = 0.500

func CornellBox(acc raytracer.Accelerator, x, y int) raytracer.Scene {

	var (
		white = &materials.DiffuseMaterial{
			Reflectance: math.Color3{R: 0.708, G: 0.743, B: 0.767},
		}
		green = &materials.DiffuseMaterial{
			Reflectance: math.Color3{R: 0.115, G: 0.441, B: 0.101},
		}
		red = &materials.DiffuseMaterial{
			Reflectance: math.Color3{R: 0.651, G: 0.059, B: 0.061},
		}
	)

	//bottom
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{},
		math.Vector3{Z: s},
		math.Vector3{X: s},
		white,
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Z: s},
		math.Vector3{X: s, Z: s},
		math.Vector3{X: s},
		white,
	))

	//left wall
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{X: s},
		math.Vector3{X: s, Z: s},
		math.Vector3{X: s, Y: s},
		red,
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{X: s, Z: s},
		math.Vector3{X: s, Z: s, Y: s},
		math.Vector3{X: s, Y: s},
		red,
	))

	//right wall
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Z: s},
		math.Vector3{},

		math.Vector3{Y: s},
		green,
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Z: s, Y: s},
		math.Vector3{Z: s},
		math.Vector3{Y: s},
		green,
	))

	//back wall
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Z: s},
		math.Vector3{Y: s, Z: s},
		math.Vector3{X: s, Z: s},
		white,
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Y: s, Z: s},
		math.Vector3{X: s, Y: s, Z: s},
		math.Vector3{X: s, Z: s},
		white,
	))

	//Ceiling
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{Y: s},
		math.Vector3{X: s, Y: s},
		math.Vector3{Z: s, Y: s},
		white,
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{X: s, Y: s},
		math.Vector3{X: s, Y: s, Z: s},
		math.Vector3{Z: s, Y: s},
		white,
	))

	camera := raytracer.GetCamera(
		math.Vector3{X: s / 2, Y: s / 2, Z: -s * 3 / 2},
		math.Vector3{Y: 1.0},
		math.Vector3{Z: 1.0},
		45.,
		1.,
		x,
		y,
	)
	var light = lights.PointLight{
		Color:  &math.Color3{R: 1.0, G: 1.0, B: 1.0},
		Origin: &math.Vector3{X: s / 2, Y: s - 0.01, Z: s / 2},
	}

	return raytracer.Scene{
		Camera:      &camera,
		Accelerator: acc,
		Lights:      []raytracer.Light{light},
	}
}
