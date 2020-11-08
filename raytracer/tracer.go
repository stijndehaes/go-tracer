package raytracer

import (
	"fmt"
	"github.com/stijndehaes/go-tracer/utils"
	"runtime"
	"sync"
)

type Tracer struct {
	Scene      *Scene
	Integrator Integrator
	pixels     chan Pixel
	wg         sync.WaitGroup
}

type Pixel struct {
	X, Y int
}

func (t *Tracer) RenderOnCanvas(canvas *utils.Canvas) {
	t.pixels = make(chan Pixel, 100)
	t.startWorkers(canvas)
	for x := 0; x < canvas.Size.X; x++ {
		for y := 0; y < canvas.Size.Y; y++ {
			t.pixels <- Pixel{X: x, Y: y}
		}
	}
	close(t.pixels)
	t.wg.Wait()
}

func (t *Tracer) startWorkers(canvas *utils.Canvas) {
	fmt.Printf("Starting op %d workers\n", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		t.wg.Add(1)
		go func() {
			defer t.wg.Done()
			for pixel := range t.pixels {
				ray := t.Scene.Camera.GenerateRay(pixel.X, pixel.Y)
				canvas.Put(pixel.X, pixel.Y, t.Integrator.RenderRay(&ray))
			}
		}()
	}
}
