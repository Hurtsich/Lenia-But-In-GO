package sensor

import (
	"Lenia/cell"
	"Lenia/helper"
	"Lenia/plane"
	"image"
)

type Donut struct {
	origin              image.Point
	neighbors           []*cell.Cell
	weightedNeighbors   []*cell.Cell
	r                   float64
	offset              float64
	nbNeighbors         int
	nbWeightedNeighbors int
}

func NewDonut(origin image.Point) *Donut {
	return &Donut{
		origin:              origin,
		r:                   6.00,
		offset:              2.00,
		nbNeighbors:         0,
		nbWeightedNeighbors: 0,
	}
}

func (s *Donut) Sense() float64 {
	cpt := 0.00
	cptWeighted := 0.00
	for _, blob := range s.neighbors {
		cpt += blob.GetStatus()
	}
	for _, blob := range s.weightedNeighbors {
		cptWeighted += blob.GetStatus() * 0.5
	}

	//fmt.Printf("Outer sum : %f, nb neighbors: %d = %f\n", cpt, s.nbNeighbors, cpt/float64(s.nbNeighbors))
	//fmt.Printf("Inner sum : %f, nb neighbors: %d = %f\n", cptWeighted, s.nbWeightedNeighbors, cptWeighted/float64(s.nbWeightedNeighbors))
	return (cpt / float64(s.nbNeighbors)) + (cptWeighted / float64(s.nbWeightedNeighbors))
}

func (s *Donut) Handshake(m [][]*cell.Cell) {
	var result []*cell.Cell
	var weightedResult []*cell.Cell

	xMin := s.origin.X - int(s.r*2)
	yMin := s.origin.Y - int(s.r*2)

	xMax := s.origin.X + int(s.r*2)
	yMax := s.origin.Y + int(s.r*2)

	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			distance := plane.GetDistance(s.origin, image.Point{x, y})
			if (distance >= s.r-(s.offset*2) && distance < s.r-s.offset) || (distance > s.r+2 && distance <= s.r+(s.offset*2)) {
				weightedResult = append(weightedResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbWeightedNeighbors++
			} else if distance >= s.r-s.offset && distance <= s.r+s.offset {
				result = append(result, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbNeighbors++
			}
		}
	}
	s.neighbors = result
	s.weightedNeighbors = weightedResult
}
