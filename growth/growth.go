package growth

import (
	"Lenia/helper"
	"math"
)

func DefaultGrowth(sum float64) float64 {
	µ := 7.00
	s := 4.00
	x := helper.Sqr(sum-µ) / (2 * helper.Sqr(s))
	return (2 * math.Exp(-x)) - 1
}
