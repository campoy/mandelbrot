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

package main

import (
	"errors"
	"flag"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/campoy/mandelbrot/mandelbrot"
)

var (
	output  = flag.String("out", "mandelbrot.png", "name of the output image file")
	height  = flag.Int("h", 2048, "height of the output image in pixels")
	width   = flag.Int("w", 2048, "width of the output image in pixels")
	mode    = flag.String("mode", "seq", "mode: seq, px, row, workers")
	workers = flag.Int("workers", 1, "number of workers to use")
)

func main() {
	flag.Parse()

	// open a new file
	f, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}

	// create the image
	img := mandelbrot.Create(*height, *width, mandelbrot.Mode(*mode), *workers)

	// and encoding it
	fmt := filepath.Ext(*output)
	switch fmt {
	case ".png":
		err = png.Encode(f, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(f, img, nil)
	case ".gif":
		err = gif.Encode(f, img, nil)
	default:
		err = errors.New("unkwnown format " + fmt)
	}
	// unless you can't
	if err != nil {
		log.Fatal(err)
	}
}
