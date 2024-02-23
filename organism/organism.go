package organism

import (
	"Lenia/cell"
	"Lenia/growth"
	"Lenia/sensor"
	"context"
	"fmt"
	"image"
	"math/rand"
	"sync"
)

type Organism struct {
	matrice [][]*cell.Cell
	tickers []chan float64
	starts  []chan struct{}
	signals []chan struct{}
	ctx     context.Context
}

var (
	organism Organism
)

func NewOrganism(length int) *Organism {
	palette = PurpleToYellowPalette()
	organism = Organism{matrice: make([][]*cell.Cell, length), ctx: context.Background()}
	for x := range organism.matrice {
		organism.matrice[x] = make([]*cell.Cell, length)
		fmt.Println("...")
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			fmt.Printf("Generating cell at %d, %d\n", x, y)
			if organism.matrice[x][y] == nil {
				blob := cell.NewCell()
				organism.matrice[x][y] = blob
			}
		}
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			if organism.matrice[x][y] != nil {
				blob := organism.matrice[x][y]
				filter := sensor.NewMultiCircle(image.Point{x, y})
				filter.Handshake(organism.matrice)
				blob.SetFilter(filter)
				blob.SetGrowth(growth.DefaultGrowth)
				core := blob.GetCore()
				organism.tickers = append(organism.tickers, core.GetDuration())
				organism.starts = append(organism.starts, core.GetTick())
				organism.signals = append(organism.signals, core.GetAntenna())
				go blob.Live(organism.ctx)
			}
		}
		fmt.Printf("Init col number %d\n", x)
	}
	setupMultiCircleLenia(organism.matrice)
	return &organism
}

func (o *Organism) Breathe(duration float64) {
	wg := new(sync.WaitGroup)
	batch := 500

	for i := 0; i < len(o.tickers); i += batch {
		j := i + batch
		if j > len(o.tickers) {
			j = len(o.tickers)
		}
		wg.Add(1)
		go startDuration(wg, o.tickers[i:j], duration)
	}
	wg.Wait()
	for i := 0; i < len(o.signals); i += batch {
		j := i + batch
		if j > len(o.signals) {
			j = len(o.signals)
		}
		wg.Add(1)
		go kickStartLife(wg, o.signals[i:j])
	}
	wg.Wait()
	for i := 0; i < len(o.starts); i += batch {
		j := i + batch
		if j > len(o.starts) {
			j = len(o.starts)
		}
		wg.Add(1)
		go tellThem(wg, o.starts[i:j])
	}
	wg.Wait()
}

func startDuration(wg *sync.WaitGroup, ticks []chan float64, duration float64) {
	defer wg.Done()
	for _, tick := range ticks {
		tick <- duration
	}
}

func kickStartLife(wg *sync.WaitGroup, starts []chan struct{}) {
	defer wg.Done()
	for _, start := range starts {
		<-start
	}
}

func tellThem(wg *sync.WaitGroup, starts []chan struct{}) {
	defer wg.Done()
	for _, start := range starts {
		start <- struct{}{}
	}
}
func (o *Organism) Die() {
	o.ctx.Done()
}

func randomStatus() float64 {
	f := rand.Float64()
	if f > 0.75 {
		return f
	}
	return 0.00
}
