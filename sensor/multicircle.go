package sensor

import (
	"Lenia/cell"
	"Lenia/helper"
	"Lenia/plane"
	"image"
)

type MultiCircle struct {
	origin            image.Point
	neighbors         []*cell.Cell
	weightedNeighbors []*cell.Cell
	r                 float64
}

func NewMultiCircle(origin image.Point) *MultiCircle {
	return &MultiCircle{
		origin: origin,
		r:      5.00,
	}
}

func (s *MultiCircle) Sense() float64 {
	cpt := 0.00
	for _, blob := range s.neighbors {
		cpt += blob.GetStatus()
	}
	for _, blob := range s.weightedNeighbors {
		cpt += blob.GetStatus() * 0.5
	}

	return cpt
}

func (s *MultiCircle) Handshake(m [][]*cell.Cell) {
	var result []*cell.Cell
	var weightedResult []*cell.Cell

	for x, col := range m {
		for y, _ := range col {
			distance := plane.GetDistance(s.origin, image.Point{helper.Mod(x, len(m[0])), helper.Mod(y, len(m))})
			if (distance >= s.r-3 && distance <= s.r-2) || (distance >= s.r+2 && distance <= s.r+3) {
				weightedResult = append(weightedResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
			} else if distance == s.r {
				result = append(result, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
			}
		}
	}
	s.neighbors = result
	//fmt.Printf("I have %d outerNeighbors\n", len(result))
	s.weightedNeighbors = weightedResult
	//fmt.Printf("I have %d weighted outerNeighbors\n", len(weightedResult))
}
