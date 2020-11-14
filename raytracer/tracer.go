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
	pixels     chan PixelRow
	wg         sync.WaitGroup
}

type PixelRow struct {
	X int
}

func (t *Tracer) RenderOnCanvas(canvas *utils.Canvas) {
	t.pixels = make(chan PixelRow, 100)
	t.startWorkers(canvas)
	for x := 0; x < canvas.Size.X; x++ {
		t.pixels <- PixelRow{X: x}
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
			for pixelRow := range t.pixels {
				for y := 0; y < canvas.Size.Y; y++ {
					ray := t.Scene.Camera.GenerateRay(pixelRow.X, y)
					canvas.Put(pixelRow.X, y, t.Integrator.RenderRay(&ray))
				}

			}
		}()
	}
}
