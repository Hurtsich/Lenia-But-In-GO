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
		r:                25.00,
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
	//fmt.Printf("Inner sum : %f, nb neighbors: %d = %f\n", cptInner, s.nbInnerNeighbors, cptInner/float64(s.nbInnerNeighbors))
	return cptOuter / float64(s.nbOuterNeighbors), cptInner / float64(s.nbInnerNeighbors)
}

func (s *Donut) Handshake(m [][]*cell.Cell) {
	var outerResult []*cell.Cell
	var innerResult []*cell.Cell

	xMin := s.origin.X - int(s.r*2)
	yMin := s.origin.Y - int(s.r*2)

	xMax := s.origin.X + int(s.r*2)
	yMax := s.origin.Y + int(s.r*2)

	for x := xMin; x < xMax; x++ {
		for y := yMin; y < yMax; y++ {
			distance := plane.GetDistance(s.origin, image.Point{x, y})
			if distance <= s.r && distance > 22.5 {
				outerResult = append(outerResult, m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))])
				s.nbOuterNeighbors++
			} else if distance <= 22.5 {
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
