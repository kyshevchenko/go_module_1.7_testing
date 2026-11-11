package stringutil

import (
	"errors"
	"testing"
)

// TestReverse — табличный тест Reverse: разные классы строк.
func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"single char", "a", "a"},
		{"plain", "hello", "olleh"},
		{"palindrome", "level", "level"},
		{"unicode", "héllo", "olléh"},
		{"cyrillic", "привет", "тевирп"},
		{"emoji", "a🎉b", "b🎉a"},
		{"spaces", "ab cd", "dc ba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Reverse(tt.input); got != tt.want {
				t.Errorf("Reverse(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestReverseInvariant — инвариант: двойной разворот возвращает исходную строку.
func TestReverseInvariant(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"single char", "x"},
		{"plain", "hello, world"},
		{"palindrome", "racecar"},
		{"unicode", "héllo wörld"},
		{"cyrillic", "привет мир"},
		{"emoji", "🎉🚀 Go!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Reverse(Reverse(tt.input)); got != tt.input {
				t.Errorf("Reverse(Reverse(%q)) = %q, want %q", tt.input, got, tt.input)
			}
		})
	}
}

// TestToUpper — табличный тест ToUpper (пока только латиница в нижнем регистре).
func TestToUpper(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"single char", "a", "A"},
		{"lowercase", "hello", "HELLO"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := ToUpper(tt.input); got != tt.want {
				t.Errorf("ToUpper(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

// TestAdd — табличный тест Add: положительные, отрицательные, нули.
func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -2, -3, -5},
		{"mixed", -2, 3, 1},
		{"zero", 0, 0, 0},
		{"big numbers", 1_000_000, 2_000_000, 3_000_000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

// TestDivide — табличный тест Divide: успех и ошибка деления на ноль.
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
		err      error
	}{
		{"divide evenly", 6, 3, 2, nil},
		{"negative result", -6, 3, -2, nil},
		{"integer truncation", 7, 2, 3, nil},
		{"divide by one", 42, 1, 42, nil},
		{"divide by zero", 1, 0, 0, ErrDivisionByZero},
		{"zero divided", 0, 5, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Divide(tt.a, tt.b)

			if !errors.Is(err, tt.err) {
				t.Errorf("Divide(%d, %d) error = %v, want %v", tt.a, tt.b, err, tt.err)
			}
			if got != tt.expected {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
