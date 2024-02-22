package helper

func Mod(a, b int) int {
	return (a%b + b) % b
}

func Abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func Sqr(a float64) float64 {
	return a * a
}
