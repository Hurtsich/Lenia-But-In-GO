package plane

import (
	"Lenia/helper"
	"image"
	"math"
)

func GetDistance(sp image.Point, tp image.Point) float64 {
	result := math.Sqrt(helper.Sqr(float64(sp.X-tp.X)) + helper.Sqr(float64(sp.Y-tp.Y)))
	//fmt.Printf("Get distance between %v and %v : %f\n", sp, tp, result)
	return result
}

func PointEqual(p1, p2 image.Point) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
