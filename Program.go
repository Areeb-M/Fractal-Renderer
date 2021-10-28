package main

import (
	"flag"
	"image/png"
	"os"
)

func main() {
	fractalFunctionArgument := flag.String("function", "mandelbrot", "Name of the fractal equation to render")
	renderXArgument := flag.Float64("rx", -0.75, "X coordinate for the center of the render")
	renderIArgument := flag.Float64("ri", 0, "I coordinate for the center of the render")
	renderWidthArgument := flag.Int("width", 1600, "The width of the render in pixels")
	renderHeightArgument := flag.Int("height", 900, "The height of the render in pixels")
	renderZoomArgument := flag.Float64("zoom", 1.0, "The zoom level of the render")
	renderOutputArgument := flag.String("output", "output.png", "The filename for the output")
	fractalIterationsArgument := flag.Int("iterations", 1024, "The maximum number of iterations for each point being rendered")

	flag.Parse()

	var fractalFunction FractalFunction

	if *fractalFunctionArgument == "mandelbrot" {
		fractalFunction = MandelbrotSet
	}

	params := ProduceRasterParameters(*renderXArgument, *renderIArgument, *renderWidthArgument,
		*renderHeightArgument, *renderZoomArgument, *fractalIterationsArgument, 2.0, fractalFunction)

	img := Engine(params)

	f, _ := os.Create(*renderOutputArgument)
	defer f.Close()

	png.Encode(f, img.SubImage(img.Rect))
}

func ProduceRasterParameters(x float64, i float64, width int, height int, zoom float64, maxIterations int, divergenceThreshold float64, fractalFunction FractalFunction) RasterParameters {
	// This scale was chosen arbitrarily to give good results at zoom = 1
	var scale float64 = 3.2

	// The size of each pixel in units
	delta := scale / float64(width) / zoom

	return RasterParameters{
		fractalFunction,
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
