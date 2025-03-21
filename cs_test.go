package cosine_similarity

import (
	"testing"
)

func Test_CS_Float64(t *testing.T) {
	var tt = []struct {
		a []float64
		b []float64
		e float64
	}{
		{[]float64{}, []float64{}, 0},
		{[]float64{0, 1, 0, 1, 0, 1, 1, 0}, []float64{0, 0, 0, 1, 1, 1, 1, 0}, 0.75},
		{[]float64{1, 1, 0, 0, 1}, []float64{1, 1, 0, 0, 1}, 1},
		{[]float64{0, 1, 0, 1, 1}, []float64{1, 0, 1, 0, 0}, 0},
		{[]float64{1, 1, 0}, []float64{1, 0, 1}, 0.5},
		{[]float64{0, 1, 1, 1, 0}, []float64{1, 0}, 0},
	}
	for _, tc := range tt {
		actual := CosineSimilarity(tc.a, tc.b)
		if actual != tc.e {
			t.Errorf("CS(%v, %v) = %v, want %v", tc.a, tc.b, actual, tc.e)
		}
	}
}

func Test_CS_Float32(t *testing.T) {
	var tt = []struct {
		a []float32
		b []float32
		e float64
	}{
		{[]float32{}, []float32{}, 0},
		{[]float32{0, 1, 0, 1, 0, 1, 1, 0}, []float32{0, 0, 0, 1, 1, 1, 1, 0}, 0.75},
		{[]float32{1, 1, 0, 0, 1}, []float32{1, 1, 0, 0, 1}, 1},
		{[]float32{0, 1, 0, 1, 1}, []float32{1, 0, 1, 0, 0}, 0},
		{[]float32{1, 1, 0}, []float32{1, 0, 1}, 0.5},
		{[]float32{0, 1, 1, 1, 0}, []float32{1, 0}, 0},
	}
	for _, tc := range tt {
		actual := CosineSimilarity(tc.a, tc.b)
		if actual != tc.e {
			t.Errorf("CS(%v, %v) = %v, want %v", tc.a, tc.b, actual, tc.e)
		}
	}
}
