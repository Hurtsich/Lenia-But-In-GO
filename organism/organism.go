package organism

import (
	"Lenia/cell"
	"Lenia/growth"
	"Lenia/helper"
	"Lenia/sensor"
	"context"
	"fmt"
	"image"
	"math/rand"
)

type Organism struct {
	matrice [][]*cell.Cell
	tickers []chan float64
	starts  []chan struct{}
	ctx     context.Context
}

var (
	organism Organism
)

func NewOrganism(length int) *Organism {
	organism = Organism{matrice: make([][]*cell.Cell, length), ctx: context.Background()}
	for x := range organism.matrice {
		organism.matrice[x] = make([]*cell.Cell, length)
		fmt.Println("...")
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			fmt.Printf("Generating cell at %d, %d\n", x, y)
			if organism.matrice[x][y] == nil {
				blob := cell.NewCell(randomStatus())
				organism.matrice[x][y] = blob
			}
		}
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			if organism.matrice[x][y] != nil {
				blob := organism.matrice[x][y]
				filter := sensor.NewDonut(image.Point{x, y})
				filter.Handshake(organism.matrice)
				blob.SetFilter(filter)
				blob.SetGrowth(growth.SmoothGrowth)
				organism.tickers = append(organism.tickers, blob.GetDuration())
				organism.starts = append(organism.starts, blob.GetTick())
				go blob.Live(organism.ctx)
			}
		}
	}
	return &organism
}

func (o *Organism) Breathe(duration float64) {
	for _, ticker := range o.tickers {
		ticker <- duration
	}
	for _, start := range o.starts {
		start <- struct{}{}
	}
}

func (o *Organism) Die() {
	o.ctx.Done()
}

func (o *Organism) Photo() *image.Paletted {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{len(o.matrice[0]), len(o.matrice)}
	photo := image.NewPaletted(image.Rectangle{topLeft, bottomRight}, helper.GetPalette())
	randx := rand.Intn(len(o.matrice))
	randy := rand.Intn(len(o.matrice[0]))
	fmt.Printf("%d:%d = %f\n", randx, randy, o.matrice[randx][randy].GetStatus())
	for col, cellColumn := range o.matrice {
		for row, blob := range cellColumn {
			photo.Set(col, row, helper.GetColor(blob.GetStatus()))
		}
	}
	return photo
}

func randomStatus() float64 {
	f := rand.Float64()
	if f > 0.75 {
		return 1.00
	}
	return 0.00
}
