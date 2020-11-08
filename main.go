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
		&math.Vector3{X: -8},
		&math.Vector3{Y: 1.0},
		&math.Vector3{X: 1.0},
		75.0,
		1.0,
		1000,
		1000,
	)

	sphere1 := geometry.Sphere{
		Radius:         1.0,
		SphereMaterial: &materials.DiffuseMaterial{Reflectance: &math.Color3{R: gomath.Pi}},
	}

	sphere2 := geometry.Sphere{
		Center:         math.Vector3{Z: -3.0},
		Radius:         1.0,
		SphereMaterial: &materials.DiffuseMaterial{Reflectance: &math.Color3{G: gomath.Pi}},
	}

	sphere3 := geometry.Sphere{
		Center:         math.Vector3{Z: 3.0},
		Radius:         1.0,
		SphereMaterial: &materials.DiffuseMaterial{Reflectance: &math.Color3{B: gomath.Pi}},
	}
	acc := accelerators.DumbAccelerator{}

	acc.AddGeometry(&sphere1)
	acc.AddGeometry(&sphere2)
	acc.AddGeometry(&sphere3)

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

	for x := 0; x < canvas.Size.X; x++ {
		for y := 0; y < canvas.Size.Y; y++ {
			ray := camera.GenerateRay(x, y)
			canvas.Put(x, y, integrator.RenderRay(&ray))
		}
	}
	err := canvas.SaveAsFile("out.png")
	if err != nil {
		panic(err)
	}
}
