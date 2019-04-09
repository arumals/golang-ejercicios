package longstring

import "testing"

func BenchmarkMedianteConcatenacion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MedianteConcatenacion(100)
	}
}

func BenchmarkMedianteArreglo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MedianteArreglo(100)
	}
}

func BenchmarkMedianteBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MedianteBuffer(100)
	}
}
