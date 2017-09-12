// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mandelbrot

import (
	"image"
	"image/color"
	"sync"
)

type img struct {
	h, w int
	m    [][]color.RGBA
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }

type Mode string

const (
	Sequential Mode = "seq"
	Pixel      Mode = "px"
	Row        Mode = "row"
	Workers    Mode = "workers"
)

func Create(height, width int, mode Mode, workers int) image.Image {
	data := make([][]color.RGBA, height)
	for i := range data {
		data[i] = make([]color.RGBA, width)
	}

	m := &img{height, width, data}

	switch mode {
	case Sequential:
		seqFillImg(m)
	case Pixel:
		oneToOneFillImg(m)
	case Row:
		onePerRowFillImg(m)
	case Workers:
		nWorkersFillImg(m, workers)
	default:
		panic("unknown mode")
	}

	return m
}

// sequential
// time mandelbrot -out=out.png -h=4096 -w=4096
// real	0m24.381s
// user	0m24.052s
// sys	0m0.174s
func seqFillImg(m *img) {
	for i, row := range m.m {
		for j := range row {
			fillPixel(m, i, j)
		}
	}
}

// one goroutine per pixel
// time mandelbrot -out=out.png -h=4096 -w=4096
// real	0m16.740s
// user	0m40.663s
// sys	0m2.186s
func oneToOneFillImg(m *img) {
	var wg sync.WaitGroup
	wg.Add(m.h * m.w)
	for i, row := range m.m {
		for j := range row {
			go func(i, j int) {
				fillPixel(m, i, j)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
}

// one per row of pixels
// time mandelbrot -out=out.png -h=4096 -w=4096
// real	0m12.156s
// user	0m27.838s
// sys	0m0.209s
func onePerRowFillImg(m *img) {
	var wg sync.WaitGroup
	wg.Add(m.h)
	for i := range m.m {
		go func(i int) {
			for j := range m.m[i] {
				fillPixel(m, i, j)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 4 workers per CPU
// real	0m17.304s
// user	0m40.615s
// sys	0m2.517s
func nWorkersFillImg(m *img, workers int) {
	var wg sync.WaitGroup
	wg.Add(workers)

	c := make(chan struct{ i, j int }, m.h*m.w)
	for i := 0; i < workers; i++ {
		go func() {
			for t := range c {
				fillPixel(m, t.i, t.j)
			}
			wg.Done()
		}()
	}

	for i, row := range m.m {
		for j := range row {
			c <- struct{ i, j int }{i, j}
		}
	}
	close(c)
	wg.Wait()
}

func fillPixel(m *img, i, j int) {
	// normalized from -2.5 to 1
	xi := 3.5*float64(i)/float64(m.w) - 2.5
	// normalized from -1 to 1
	yi := 2*float64(j)/float64(m.h) - 1

	const maxI = 1000
	x, y := 0., 0.
	for i := 0; (x*x+y*y < 4) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	paint(&m.m[i][j], x, y)
}

func paint(c *color.RGBA, x, y float64) {
	n := byte(x * y)
	c.R, c.G, c.B, c.A = n, n, n, 255
}
