package growth

func SmoothGrowth(outerSum float64, innerSum float64) float64 {
	if innerSum >= 0.5 && (outerSum >= 0.26 && outerSum <= 0.46) {
		return 1.00
	} else if innerSum < 0.5 && (outerSum >= 0.27 && outerSum <= 0.36) {
		return 1.00
	}
	return 0.00
}
