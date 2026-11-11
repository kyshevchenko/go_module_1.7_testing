package stringutil

import (
	"errors"
	"testing"
)

// TestReverse — белый ящик: Reverse переворачивает строку.
func TestReverse(t *testing.T) {
	if got := Reverse("hello"); got != "olleh" {
		t.Errorf("Reverse(\"hello\") = %q, want \"olleh\"", got)
	}
	if got := Reverse(""); got != "" {
		t.Errorf("Reverse(\"\") = %q, want \"\"", got)
	}
	if got := Reverse("a"); got != "a" {
		t.Errorf("Reverse(\"a\") = %q, want \"a\"", got)
	}
}

// TestToUpper — белый ящик: только латиница в нижнем регистре.
// Нелатинские символы ToUpper пока не поддерживает — исправим позже.
func TestToUpper(t *testing.T) {
	if got := ToUpper("hello"); got != "HELLO" {
		t.Errorf("ToUpper(\"hello\") = %q, want \"HELLO\"", got)
	}
	if got := ToUpper(""); got != "" {
		t.Errorf("ToUpper(\"\") = %q, want \"\"", got)
	}
}

// TestAdd — белый ящик: сложение целых чисел.
func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2, 3) = %d, want 5", got)
	}
	if got := Add(-2, 3); got != 1 {
		t.Errorf("Add(-2, 3) = %d, want 1", got)
	}
	if got := Add(0, 0); got != 0 {
		t.Errorf("Add(0, 0) = %d, want 0", got)
	}
}

// TestDivide — белый ящик: успешное деление и деление на ноль.
func TestDivide(t *testing.T) {
	got, err := Divide(6, 3)
	if err != nil {
		t.Fatalf("Divide(6, 3) unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("Divide(6, 3) = %d, want 2", got)
	}

	_, err = Divide(1, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Divide(1, 0) error = %v, want %v", err, ErrDivisionByZero)
	}
}
