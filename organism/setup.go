package organism

import (
	"Lenia/cell"
	"Lenia/collection"
	"Lenia/helper"
	"Lenia/plane"
	"image"
	"math/rand"
)

func setupSmoothLife(m [][]*cell.Cell) {
	randPoints := getRandPoints(len(m))

	for x, col := range m {
		for y, blob := range col {
			if collection.Contains(randPoints, image.Point{x, y}, plane.PointEqual) {
				origin := image.Point{x, y}
				r := 25.00
				xMin := origin.X - int(r*2)
				yMin := origin.Y - int(r*2)

				xMax := origin.X + int(r*2)
				yMax := origin.Y + int(r*2)

				for x := xMin; x < xMax; x++ {
					for y := yMin; y < yMax; y++ {
						distance := plane.GetDistance(origin, image.Point{x, y})
						if distance <= r && distance > 22.5 {
							m[helper.Mod(x, len(m[0]))][helper.Mod(y, len(m))].SetStatus(1.00)
						}
					}
				}
			} else {
				if blob.GetStatus() != 1 {
					blob.SetStatus(0.00)
				}
			}
		}
	}
}

func getRandPoints(length int) []image.Point {
	var result []image.Point
	for i := 0; i < 40; i++ {
		result = append(result, image.Point{X: rand.Intn(length), Y: rand.Intn(length)})
	}
	return result
}
