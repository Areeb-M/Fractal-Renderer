package main

import (
	"image/color"
)

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

	r := int(colorA.R) + int(float32(colorB.R-colorA.R)*colorLerp)
	g := int(colorA.G) + int(float32(colorB.G-colorA.G)*colorLerp)
	b := int(colorA.B) + int(float32(colorB.B-colorA.B)*colorLerp)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

// Red, Orange, Yellow, Green, Blue, Indigo, Violet
var rainbowPalette = []color.RGBA{{255, 0, 0, 255}, {255, 165, 0, 255}, {255, 255, 0, 255}, {0, 255, 0, 255},
	{0, 0, 255, 255}, {75, 0, 130, 255}, {143, 0, 255, 255}}

var rainbowPalette2 = []color.RGBA{{0, 0, 0, 255}, {0, 255, 0, 255}, {0, 255, 255, 255}, {255, 255, 255, 255}}
