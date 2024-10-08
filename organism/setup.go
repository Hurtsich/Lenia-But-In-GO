package organism

import (
	"Lenia/cell"
	"Lenia/collection"
	"Lenia/helper"
	"Lenia/plane"
	"image"
	"math/rand"
)

func setupSimpleLenia(m [][]*cell.Cell) {
	r := 15.00
	offset := 2.00
	randPoints := getRandPoints(r, offset, len(m))

	for x, col := range m {
		for y, blob := range col {
			if collection.Contains(randPoints, image.Point{x, y}, plane.PointEqual) {
				origin := image.Point{x, y}
				xMin := origin.X - int(r*2)
				yMin := origin.Y - int(r*2)

				xMax := origin.X + int(r*2)
				yMax := origin.Y + int(r*2)

				for x := xMin; x < xMax; x++ {
					for y := yMin; y < yMax; y++ {
						distance := plane.GetDistance(origin, image.Point{x, y})
						if (distance >= r-(offset*2) && distance < r-offset) || (distance > r+offset && distance <= r+(offset*2)) {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(0.2)
						} else if distance > r-offset && distance < r+offset {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(0.5)
						} else if distance == r {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(1.00)
						}
					}
				}
			} else {
				if blob.GetStatus() != 0.00 {
					continue
				}
			}
		}
	}
}

func setupMultiCircleLenia(m [][]*cell.Cell) {
	r := 35.00
	offset := 9.00
	randPoints := getRandPoints(r, offset, len(m))

	for x, col := range m {
		for y, blob := range col {
			if collection.Contains(randPoints, image.Point{x, y}, plane.PointEqual) {
				origin := image.Point{x, y}
				xMin := origin.X - int(r*2)
				yMin := origin.Y - int(r*2)

				xMax := origin.X + int(r*2)
				yMax := origin.Y + int(r*2)

				for x := xMin; x < xMax; x++ {
					for y := yMin; y < yMax; y++ {
						distance := plane.GetDistance(origin, image.Point{x, y})
						if distance >= r+offset && distance <= r+offset+2 {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(0.25)
						} else if distance >= r-1 && distance <= r+1 {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(1.00)
						} else if distance >= r-(offset+2) && distance <= r-offset {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(0.5)
						}
					}
				}
			} else {
				if blob.GetStatus() != 0.00 {
					continue
				}
			}
		}
	}
}

func getRandPoints(r, offset float64, length int) []image.Point {
	var result []image.Point
	randX := rand.Intn(length)
	randY := rand.Intn(length)
	point := image.Point{X: randX, Y: randY}
	result = append(result, point)
	cpt := 0
	for cpt < 1 {
		randX = rand.Intn(length)
		randY = rand.Intn(length)
		randPoint := image.Point{randX, randY}
		distance := plane.GetDistance(point, randPoint)
		if distance <= r*2 && distance > r+offset {
			result = append(result, randPoint)
			cpt++
		}
	}
	return result
}
