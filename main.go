package main

import (
	"github.com/Steniaz/go-tracer/utils"
	"image"
	"github.com/Steniaz/go-tracer/math"
	"github.com/Steniaz/go-tracer/raytracer"
	"github.com/Steniaz/go-tracer/geometry"
	"github.com/Steniaz/go-tracer/accelerators"
	"github.com/Steniaz/go-tracer/materials"
	"github.com/Steniaz/go-tracer/integrators"
	"github.com/Steniaz/go-tracer/lights"
	gomath "math"
)

func main() {
	canvas := utils.CreateCanvas(image.Point{1000, 1000})
	camera := raytracer.GetCamera(
		&math.Vector3{X: -8},
		&math.Vector3{Y: 1.0},
		&math.Vector3{X: 1.0},
		75.0,
		1.0,
		1000,
		1000,
	)

	sphere1 :=  geometry.Sphere{
		math.Vector3{},
		1.0,
		&materials.DiffuseMaterial{&math.Color3{R: gomath.Pi}},
	}

	sphere2 :=  geometry.Sphere{
		math.Vector3{Z: -3.0},
		1.0,
		&materials.DiffuseMaterial{&math.Color3{G: gomath.Pi}},
	}

	sphere3 :=  geometry.Sphere{
		math.Vector3{Z: 3.0},
		1.0,
		&materials.DiffuseMaterial{&math.Color3{B: gomath.Pi}},
	}
	acc := accelerators.DumbAccelerator{}

	acc.AddGeometry(&sphere1)
	acc.AddGeometry(&sphere2)
	acc.AddGeometry(&sphere3)

	light := lights.PointLight{
		&math.Color3{100.0, 100.0, 100.0},
		&math.Vector3{0.0, 10.0, 0.0},
	}
	sceneLights := []raytracer.Light{light}

	scene := raytracer.Scene{
		&camera,
		&acc,
		sceneLights,
	}
	integrator := integrators.DirectLightningIntegrator{
		&scene,
	}

	for x := 0; x < canvas.Size.X; x++ {
		for y := 0; y < canvas.Size.Y; y++ {
			ray := camera.GenerateRay(x, y)
			canvas.Put(x, y, integrator.RenderRay(&ray))
		}
	}
	canvas.SaveAsFile("out.png")
}