package stringutil_test

// Чёрный ящик: тестируем только публичный API пакета stringutil.
// Внутри этого файла НЕ доступны неэкспортированные идентификаторы
// (например, внутренние переменные или функции в нижнем регистре).

import (
	"errors"
	"testing"

	"stringutil"
)

// TestReverseBlackbox — чёрный ящик: поведение Reverse как внешнего API.
func TestReverseBlackbox(t *testing.T) {
	if got := stringutil.Reverse("hello"); got != "olleh" {
		t.Errorf("stringutil.Reverse(\"hello\") = %q, want \"olleh\"", got)
	}
	if got := stringutil.Reverse(""); got != "" {
		t.Errorf("stringutil.Reverse(\"\") = %q, want \"\"", got)
	}
}

// TestToUpperBlackbox — чёрный ящик: латиница в верхний регистр.
func TestToUpperBlackbox(t *testing.T) {
	if got := stringutil.ToUpper("hello"); got != "HELLO" {
		t.Errorf("stringutil.ToUpper(\"hello\") = %q, want \"HELLO\"", got)
	}
}

// TestAddBlackbox — чёрный ящик: сложение.
func TestAddBlackbox(t *testing.T) {
	if got := stringutil.Add(2, 3); got != 5 {
		t.Errorf("stringutil.Add(2, 3) = %d, want 5", got)
	}
}

// TestDivideBlackbox — чёрный ящик: деление, включая ошибку деления на ноль.
func TestDivideBlackbox(t *testing.T) {
	got, err := stringutil.Divide(6, 3, nil)
	if err != nil {
		t.Fatalf("stringutil.Divide(6, 3) unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("stringutil.Divide(6, 3) = %d, want 2", got)
	}

	_, err = stringutil.Divide(1, 0, nil)
	if !errors.Is(err, stringutil.ErrDivisionByZero) {
		t.Errorf("stringutil.Divide(1, 0) error = %v, want %v", err, stringutil.ErrDivisionByZero)
	}
}
