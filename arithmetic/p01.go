package arithmetic

import "math"

func PrimeOrNot(input int) bool {
	if input < 2 {
		return false
	}
	root := int(math.Sqrt(float64(input)))

	for j := 2; j <= root; j++ {
		if input%j == 0 {
			return false
		}
	}
	return true
}
