package main

import (
	"image/png"
	"os"
)

func main() {
	params := RasterParameters{
		MandelbrotSet,
		16000,
		9000,
		-.9,
		.9,
		-2.0,
		1.2,
		255,
		2.0,
	}

	img := Engine(params)

	f, _ := os.Create("output.png")
	defer f.Close()

	png.Encode(f, img.SubImage(img.Rect))
}
