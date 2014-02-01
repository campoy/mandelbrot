package main

import (
	"flag"
	"image/png"
	"log"
	"os"
	"runtime"

	"github.com/campoy/mandelbrot/mandelbrot"
)

var (
	output = flag.String("out", "mandelbrot", "name of the output image file")
	format = flag.String("f", "png", "format of the output image")
	height = flag.Int("h", 2048, "height of the output image in pixels")
	width  = flag.Int("w", 2048, "width of the output image in pixels")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, err := os.Create(*output + "." + *format)
	if err != nil {
		log.Fatal(err)
	}

	img := mandelbrot.Create(*height, *width)

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
