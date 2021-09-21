package main

import (
	"image"
)

type RasterParameters struct {
	fractalFunction FractalFunction

	rasterWidth  int
	rasterHeight int

	iMin float64
	iMax float64
	xMin float64
	xMax float64

	maxIterations       int
	divergenceThreshold float64
}

func Engine(parameters RasterParameters) image.RGBA {
	raster := image.NewRGBA(image.Rect(0, 0, parameters.rasterWidth, parameters.rasterHeight))

	var dx float64 = (parameters.xMax - parameters.xMin) / float64(parameters.rasterWidth)
	var di float64 = (parameters.iMax - parameters.iMin) / float64(parameters.rasterHeight)
	var c complex128

	colorBlend := ColorBlend{rainbowPalette2[:], parameters.maxIterations}

	for rx := 0; rx < parameters.rasterWidth; rx++ {
		for ry := 0; ry < parameters.rasterHeight; ry++ {
			c = complex(dx*float64(rx)+parameters.xMin, di*float64(ry)+parameters.iMin)

			depth := parameters.fractalFunction(c, parameters.divergenceThreshold, parameters.maxIterations)
			//raster.Set(rx, ry, color.RGBA{uint8(depth), 0, 0, 255})
			raster.Set(rx, ry, colorBlend.Blend(depth))
		}
	}

	return *raster
}
