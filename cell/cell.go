package cell

import (
	"context"
	"fmt"
	"sync/atomic"
)

type Sensor interface {
	Sense() float64
	Handshake([][]*Cell)
}

type Cell struct {
	sensor Sensor
	status atomic.Value
	core   *Core
	grow   func(float64) float64
}

func NewCell() *Cell {
	blob := &Cell{
		core: NewCore(),
	}
	blob.SetStatus(0.00)
	return blob
}

func (c *Cell) Live(ctx context.Context) {
	fmt.Println("Cell ready to report !")
	for {
		select {
		case <-ctx.Done():
			break
		default:

			elapsed := <-c.core.duration

			sumNeigh := c.sensor.Sense()

			c.core.antenna <- struct{}{}

			<-c.core.tick

			val := c.grow(sumNeigh) * (1 / elapsed)
			newStatus := c.GetStatus() + val
			fmt.Printf("My neighbors gave give me %f value which means i'm growing %f much and now i'm %f\n", sumNeigh, val, newStatus)
			if newStatus > 1 {
				c.status.Store(1.00)
			} else if val < 0 {
				c.status.Store(0.00)
			} else {
				c.status.Store(newStatus)
			}
		}
	}
}

func (c *Cell) SetFilter(s Sensor) {
	c.sensor = s
}

func (c *Cell) SetGrowth(f func(float64) float64) {
	c.grow = f
}

func (c *Cell) GetStatus() float64 {
	return c.status.Load().(float64)
}

func (c *Cell) GetCore() *Core {
	return c.core
}

func (c *Cell) SetStatus(f float64) {
	c.status.Store(f)
}
