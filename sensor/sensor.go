package sensor

import (
	"Lenia/cell"
	"Lenia/helper"
	"Lenia/plane"
	"image"
)

type Donut struct {
	origin      image.Point
	neighbors   []*cell.Cell
	neighborsNb int
}

func NewDonut(origin image.Point) *Donut {
	return &Donut{
		origin: origin,
	}
}

func (s *Donut) GetNeighbors() int {
	return s.neighborsNb
}

func (s *Donut) Sense() float64 {
	cpt := 0.00
	for _, blob := range s.neighbors {
		cpt += blob.GetStatus()
	}
	return cpt
}

func (s *Donut) Handshake(m [][]*cell.Cell) {
	var result []*cell.Cell

	xMax := s.origin.X + len(m[0])
	yMax := s.origin.Y + len(m)
	for x := s.origin.X; x < xMax; x++ {
		for y := s.origin.Y; y < yMax; y++ {
			distance := plane.GetDistance(s.origin, image.Point{helper.Mod(x, len(m[0])), helper.Mod(y, len(m))})
			if distance >= 5 && distance <= 7 {
				result = append(result, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.neighborsNb++
			}
		}
	}
	s.neighbors = result
}
