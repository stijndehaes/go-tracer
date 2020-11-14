package main

import (
	"github.com/stijndehaes/go-tracer/accelerators"
	"github.com/stijndehaes/go-tracer/integrators"
	"github.com/stijndehaes/go-tracer/raytracer"
	"github.com/stijndehaes/go-tracer/scenes"
	"github.com/stijndehaes/go-tracer/utils"
	"image"
)

func main() {
	x := 4000
	y := 4000
	canvas := utils.CreateCanvas(image.Point{X: x, Y: y})
	acc := accelerators.DumbAccelerator{}
	scene := scenes.CornellBox(&acc, x, y)
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
