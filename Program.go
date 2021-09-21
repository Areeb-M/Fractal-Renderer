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
		1024,
		2.0,
	}

	params = ProduceRasterParameters(-0.7463, 0.1103, 1920, 1080, 20000, 512, 2.0)

	//params = ProduceRasterParameters(-0.04524074130409, 0.9868162207157852, 1920, 1080, 200000000, 1024, 2.0)
	// ^^ This is a really pretty output
	params = ProduceRasterParameters(-0.04524074130409, 0.9868162207157852, 16000, 9000, 200000000, 1024, 2.0)
	//params = ProduceRasterParameters(0, 0, 1600, 900, 1, 255, 2.0)

	img := Engine(params)

	f, _ := os.Create("output.png")
	defer f.Close()

	png.Encode(f, img.SubImage(img.Rect))
}

func ProduceRasterParameters(x float64, i float64, width int, height int, zoom float64, maxIterations int, divergenceThreshold float64) RasterParameters {
	var scale float64 = 3.2

	delta := scale / float64(width) / zoom

	return RasterParameters{
		MandelbrotSet,
		width,
		height,
		i - (float64(height)/2.0)*delta,
		i + (float64(height)/2.0)*delta,
		x - (float64(width)/2.0)*delta,
		x + (float64(width)/2.0)*delta,
		maxIterations,
		divergenceThreshold,
	}
}
