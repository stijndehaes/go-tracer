package main

import (
	"github.com/stijndehaes/go-tracer/accelerators"
	"github.com/stijndehaes/go-tracer/geometry"
	"github.com/stijndehaes/go-tracer/integrators"
	"github.com/stijndehaes/go-tracer/lights"
	"github.com/stijndehaes/go-tracer/materials"
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
	"github.com/stijndehaes/go-tracer/utils"
	"image"
	gomath "math"
)

func main() {
	canvas := utils.CreateCanvas(image.Point{X: 1000, Y: 1000})
	camera := raytracer.GetCamera(
		math.Vector3{X: -8},
		math.Vector3{Y: 1.0},
		math.Vector3{X: 1.0},
		75.0,
		1.0,
		1000,
		1000,
	)

	colors := []math.Color3{
		{R: gomath.Pi},
		{G: gomath.Pi},
		{B: gomath.Pi},
	}

	acc := accelerators.DumbAccelerator{}
	for x := 0; x <= 40; x = x + 1 {
		for y := 0; y <= 40; y = y + 1 {
			sphere := geometry.Sphere{
				Radius: 0.3,
				Center: math.Vector3{
					Z: -20.0 + float64(x),
					X: -5.0 + float64(y),
					Y: -3.0,
				},
				SphereMaterial: &materials.DiffuseMaterial{Reflectance: colors[(x+y)%3]},
			}
			acc.AddGeometry(&sphere)
		}
	}
	triangleLenght := 50.
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{
			Z: -1 * triangleLenght,
			X: -1 * triangleLenght,
			Y: -3.1,
		},
		math.Vector3{
			Z: triangleLenght,
			X: -1 * triangleLenght,
			Y: -3.1,
		},
		math.Vector3{
			Z: triangleLenght,
			X: 20,
			Y: -3.1,
		},
		&materials.DiffuseMaterial{Reflectance: math.Color3{R: gomath.Pi}},
	))
	acc.AddGeometry(geometry.NewTriangle(
		math.Vector3{
			Z: -1 * triangleLenght,
			X: -1 * triangleLenght,
			Y: -3.1,
		},

		math.Vector3{
			Z: triangleLenght,
			X: 20,
			Y: -3.1,
		},
		math.Vector3{
			Z: -1 * triangleLenght,
			X: triangleLenght,
			Y: -3.1,
		},
		&materials.DiffuseMaterial{Reflectance: math.Color3{R: gomath.Pi}},
	))

	var light = lights.PointLight{
		Color:  &math.Color3{R: 100.0, G: 100.0, B: 100.0},
		Origin: &math.Vector3{Y: 10.0},
	}
	sceneLights := []raytracer.Light{light}

	scene := raytracer.Scene{
		Camera:      &camera,
		Accelerator: &acc,
		Lights:      sceneLights,
	}
	integrator := integrators.DirectLightningIntegrator{
		Scene: &scene,
	}

	t := raytracer.Tracer{
		Scene:      &scene,
		Integrator: &integrator,
	}

	t.RenderOnCanvas(canvas)
	err := canvas.SaveAsFile("out.png")
	if err != nil {
		panic(err)
	}
}
