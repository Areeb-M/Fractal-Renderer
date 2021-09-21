package main

import "image/color"

type ColorBlend struct {
	palette       []color.RGBA
	maxIterations int
}

func (colorBlend *ColorBlend) Blend(iterations int) color.RGBA {
	colorLerp := (float32(iterations) / float32(colorBlend.maxIterations+1)) * float32(len(colorBlend.palette)-1)
	colorIndex := int(colorLerp)
	colorLerp -= float32(colorIndex) // leaves just the fractional portion of the lerp
	colorA := colorBlend.palette[colorIndex]
	colorB := colorBlend.palette[colorIndex+1]

	r := colorA.R + uint8(float32(colorA.R-colorB.R)*colorLerp)
	g := colorA.G + uint8(float32(colorA.G-colorB.G)*colorLerp)
	b := colorA.B + uint8(float32(colorA.B-colorB.B)*colorLerp)

	return color.RGBA{r, g, b, 255}
}

// Red, Orange, Yellow, Green, Blue, Indigo, Violet
var rainbowPalette = [...]color.RGBA{{255, 0, 0, 255}, {255, 165, 0, 255}, {255, 255, 0, 255}, {0, 255, 0, 255},
	{0, 0, 255, 255}, {75, 0, 130, 255}, {143, 0, 255, 255}}
