package organism

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"math"
	"math/rand"
)

var (
	palette []color.Color
)

func (o *Organism) Photo() *image.Paletted {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{len(o.matrice[0]), len(o.matrice)}
	photo := image.NewPaletted(image.Rectangle{topLeft, bottomRight}, palette)
	randx := rand.Intn(len(o.matrice))
	randy := rand.Intn(len(o.matrice[0]))
	fmt.Printf("%d:%d = %f\n", randx, randy, o.matrice[randx][randy].GetStatus())
	for col, cellColumn := range o.matrice {
		for row, blob := range cellColumn {
			photo.Set(row, col, GetColor(blob.GetStatus()))
		}
	}
	return photo
}

func GetColor(gradient float64) color.Color {
	pal := palette
	index := gradient * float64(len(pal)-1)
	return pal[int(math.Floor(index))]
}

func BlackToRedPalette() []color.Color {
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

func BlackToYellowPalette() []color.Color {
	start := colorful.HSLuv(0, 0, 0.05) // Black
	end := colorful.HSLuv(60, 1, 0.75)  // Light Blue

	// Generate the gradient
	gradient := make([]color.Color, 64)
	for i := range gradient {
		h := float64(i) / float64(len(gradient)-1)
		c := start.BlendLuv(end, h).Clamped()
		gradient[i] = color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255}
	}

	return gradient
}

func BlueToYellowPalette() []color.Color {
	start := colorful.HSLuv(240, 1, 0.05) // Black
	end := colorful.HSLuv(60, 1, 0.5)     // Light Blue

	// Generate the gradient
	gradient := make([]color.Color, 256)
	for i := range gradient {
		h := float64(i) / float64(len(gradient)-1)
		c := start.BlendHcl(end, h).Clamped()
		gradient[i] = color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255}
	}

	return gradient
}

func PurpleToYellowPalette() []color.Color {
	start := colorful.HSLuv(280, 1, 0.1) // Black
	end := colorful.HSLuv(60, 1, 0.5)    // Light Blue

	// Generate the gradient
	gradient := make([]color.Color, 256)
	for i := range gradient {
		h := float64(i) / float64(len(gradient)-1)
		c := start.BlendHcl(end, h).Clamped()
		gradient[i] = color.RGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255}
	}

	return gradient
}
