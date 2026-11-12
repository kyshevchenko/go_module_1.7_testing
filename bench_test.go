package stringutil

import (
	"strings"
	"testing"
)

// BenchmarkReverse — производительность Reverse на строках разной длины.
func BenchmarkReverse(b *testing.B) {
	sizes := []struct {
		name  string
		input string
	}{
		{"short_5", "hello"},
		{"medium_1k", strings.Repeat("abcd", 250)},
		{"long_100k", strings.Repeat("abcd", 25000)},
	}

	for _, size := range sizes {
		b.Run(size.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer() // подготовка (Repeat) не должна попадать в замер
			for i := 0; i < b.N; i++ {
				Reverse(size.input)
			}
		})
	}
}

// BenchmarkAdd — сложение через оператор +.
func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(1234567, 7654321)
	}
}

// addIterative — альтернативная реализация Add через итерации.
// Демонстрационная: O(b) вместо O(1).
func addIterative(a, b int) int {
	result := a
	for i := 0; i < b; i++ {
		result++
	}
	return result
}

// BenchmarkAddIterative — сравнение с addIterative.
func BenchmarkAddIterative(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		addIterative(1234567, 7654321)
	}
}
