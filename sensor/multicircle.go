package sensor

import (
	"Lenia/cell"
	"Lenia/helper"
	"Lenia/plane"
	"image"
)

type MultiCircle struct {
	origin           image.Point
	neighbors        []*cell.Cell
	nbNeighbors      int
	innerNeighbors   []*cell.Cell
	nbInnerNeighbors int
	outerNeighbors   []*cell.Cell
	nbOuterNeighbors int
	r                float64
	offset           float64
}

func NewMultiCircle(origin image.Point) *MultiCircle {
	return &MultiCircle{
		origin: origin,
		r:      35.00,
		offset: 9.00,
	}
}

func (s *MultiCircle) Sense() float64 {
	cpt := 0.00
	cptInner := 0.00
	cptOuter := 0.00
	for _, blob := range s.neighbors {
		cpt += blob.GetStatus()
	}
	for _, blob := range s.innerNeighbors {
		cptInner += blob.GetStatus() * 0.5
	}
	for _, blob := range s.outerNeighbors {
		cptOuter += blob.GetStatus() * 0.25
	}

	//fmt.Printf("Neighbors sum : %f, nb neighbors: %d = %f\n", cpt, s.nbNeighbors, cpt/float64(s.nbNeighbors))
	//fmt.Printf("Inner sum : %f, nb neighbors: %d = %f\n", cptInner, s.nbInnerNeighbors, cptInner/float64(s.nbInnerNeighbors))
	//fmt.Printf("Outer sum : %f, nb neighbors: %d = %f\n", cptOuter, s.nbOuterNeighbors, cptOuter/float64(s.nbOuterNeighbors))
	return (cpt / float64(s.nbNeighbors)) + (cptInner / float64(s.nbInnerNeighbors)) + (cptOuter / float64(s.nbOuterNeighbors))
}

func (s *MultiCircle) Handshake(m [][]*cell.Cell) {
	var result []*cell.Cell
	var innerResult []*cell.Cell
	var outerResult []*cell.Cell

	xMin := s.origin.X - int(s.r*2)
	yMin := s.origin.Y - int(s.r*2)

	xMax := s.origin.X + int(s.r*2)
	yMax := s.origin.Y + int(s.r*2)

	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			distance := plane.GetDistance(s.origin, image.Point{x, y})
			if distance >= s.r+s.offset && distance <= s.r+s.offset+2 {
				outerResult = append(outerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbOuterNeighbors++
			} else if distance >= s.r-1 && distance <= s.r+1 {
				result = append(result, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbNeighbors++
			} else if distance >= s.r-(s.offset+2) && distance <= s.r-s.offset {
				innerResult = append(innerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbInnerNeighbors++
			}
		}
	}
	s.neighbors = result
	//fmt.Printf("I have %d neighbors\n", len(result))
	s.innerNeighbors = innerResult
	s.outerNeighbors = outerResult
	//fmt.Printf("I have %d weighted neighbors\n", len(innerResult))
}
