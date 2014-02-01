package mandelbrot

import (
	"image"
	"image/color"
)

type img struct {
	h, w int
	m    [][]color.RGBA
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }

func Create(h, w int) image.Image {
	c := make([][]color.RGBA, h)
	for i := range c {
		c[i] = make([]color.RGBA, w)
	}

	m := &img{h, w, c}
	fillImg(m)

	return m
}

func fillImg(m *img) {
	for i, row := range m.m {
		for j := range row {
			fillPixel(m, i, j)
		}
	}
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
	c.R = n
	c.G = n
	c.B = n
	c.A = 255
}
