package sensor

import (
	"Lenia/cell"
	"Lenia/helper"
	"Lenia/plane"
	"image"
)

type Donut struct {
	origin           image.Point
	outerNeighbors   []*cell.Cell
	innerNeighbors   []*cell.Cell
	r                float64
	nbInnerNeighbors int
	nbOuterNeighbors int
}

func NewDonut(origin image.Point) *Donut {
	return &Donut{
		origin:           origin,
		r:                10.00,
		nbInnerNeighbors: 0,
		nbOuterNeighbors: 0,
	}
}

func (s *Donut) Sense() (float64, float64) {
	cptInner := 0.00
	cptOuter := 0.00
	for _, blob := range s.outerNeighbors {
		cptOuter += blob.GetStatus()
	}
	for _, blob := range s.innerNeighbors {
		cptInner += blob.GetStatus()
	}
	//fmt.Printf("Outer sum : %f, nb neighbors: %d = %f\n", cptOuter, s.nbOuterNeighbors, cptOuter/float64(s.nbOuterNeighbors))
	return cptOuter / float64(s.nbOuterNeighbors), cptInner / float64(s.nbInnerNeighbors)
}

func (s *Donut) Handshake(m [][]*cell.Cell) {
	var outerResult []*cell.Cell
	var innerResult []*cell.Cell

	xMax := s.origin.X + len(m)
	yMax := s.origin.Y + len(m[0])

	for x := s.origin.X; x < xMax; x++ {
		for y := s.origin.Y; y < yMax; y++ {
			distance := plane.GetDistance(s.origin, image.Point{x, y})
			if distance <= s.r && distance > 1 {
				outerResult = append(outerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbOuterNeighbors++
			} else if distance <= 1 {
				innerResult = append(innerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbInnerNeighbors++
			}
		}
	}
	//for x, col := range m {
	//	for y, _ := range col {
	//		distance := plane.GetDistance(s.origin, image.Point{helper.Mod(x, len(m[0])), helper.Mod(y, len(m))})
	//		if distance <= s.r && distance > 1 {
	//			outerResult = append(outerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
	//			s.nbOuterNeighbors++
	//		} else if distance <= 1 {
	//			innerResult = append(innerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
	//			s.nbInnerNeighbors++
	//		}
	//	}
	//}
	s.outerNeighbors = outerResult
	//fmt.Printf("I have %d outerNeighbors\n", len(outerResult))
	s.innerNeighbors = innerResult
	//fmt.Printf("I have %d weighted outerNeighbors\n", len(innerResult))
}
