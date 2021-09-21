package main

import (
	"image/jpeg"
	"os"
)

func main() {
	params := RasterParameters{
		MandelbrotSet,
		1600,
		900,
		-.9,
		.9,
		-1.6,
		1.6,
		255,
		2.0,
	}

	img := Engine(params)

	f, _ := os.Create("output.jpg")
	defer f.Close()

	jpeg.Encode(f, img.SubImage(img.Rect), &jpeg.Options{jpeg.DefaultQuality})
}
