package cosine_similarity

import "math"

// CosineSimilarity calculates the cosign similarity of 2 vectors
// "Cosine similarity is the cosine of the angle between the vectors;
// that is, it is the dot product of the vectors divided by the product
// of their lengths"
func CosineSimilarity[T float32 | float64](a, b []T) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	a, b = norm(a, b)
	var sumA, sumB, sumAB float64
	for i, v := range a {
		sumA += float64(v) * float64(v)
		sumB += float64(b[i]) * float64(b[i])
		sumAB += float64(v) * float64(b[i])
	}
	return sumAB / (math.Sqrt(sumA * sumB))
}

func norm[T float32 | float64](a []T, b []T) ([]T, []T) {
	n := func(c []T, i int) []T {
		return append(c, make([]T, i)...)
	}
	if len(a) > len(b) {
		return a, n(b, len(a)-len(b))
	} else if len(b) > len(a) {
		return n(a, len(b)-len(a)), b
	}
	return a, b
}
