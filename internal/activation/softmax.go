package activation

import "math"

// Softmax activation function.
func Softmax(c, z, sum float64) float64 {
	return (c * math.Exp(z)) / (c * sum)
}
