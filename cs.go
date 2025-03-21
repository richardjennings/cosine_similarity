package cosine_similarity

import "math"

// CosineSimilarity calculates the cosign similarity of 2 vectors
// "Cosine similarity is the cosine of the angle between the vectors;
// that is, it is the dot product of the vectors divided by the product
// of their lengths"
func CosineSimilarity(a, b []float64) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	a, b = norm(a, b)
	var sumA, sumB, sumAB float64
	for i, v := range a {
		sumA += v * v
		sumB += b[i] * b[i]
		sumAB += v * b[i]
	}
	return sumAB / (math.Sqrt(sumA * sumB))
}

func norm(a []float64, b []float64) ([]float64, []float64) {
	n := func(c []float64, i int) []float64 {
		return append(c, make([]float64, i)...)
	}
	if len(a) > len(b) {
		return a, n(b, len(a)-len(b))
	} else if len(b) > len(a) {
		return n(a, len(b)-len(a)), b
	}
	return a, b
}
