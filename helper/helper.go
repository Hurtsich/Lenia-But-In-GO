package helper

import (
	"github.com/lucasb-eyer/go-colorful"
	"image/color"
	"math"
)

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func Sqr(a float64) float64 {
	return a * a
}

func GetColor(gradient float64) color.Color {
	pal := GetPalette()
	index := gradient * float64(len(pal)-1)
	return pal[int(math.Floor(index))]
}

func GetPalette() []color.Color {
	start := colorful.Hsl(0, 0, 0)    // Black
	end := colorful.Hsl(0.5, 1, 0.75) // Light Blue

	// Generate the gradient
	gradient := make([]color.Color, 64)
	for i := range gradient {
		h := float64(i) / float64(len(gradient)-1)
		c := start.BlendLuv(end, h).Clamped()
		gradient[i] = color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255}
	}

	return gradient
}
