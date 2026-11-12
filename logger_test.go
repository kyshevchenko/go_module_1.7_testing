package stringutil

import (
	"errors"
	"strings"
	"testing"
)

// mockLogger — ручной мок: записывает вызовы вместо вывода на экран.
type mockLogger struct {
	infoCalls  []string
	errorCalls []string
}

func (m *mockLogger) Info(msg string)  { m.infoCalls = append(m.infoCalls, msg) }
func (m *mockLogger) Error(msg string) { m.errorCalls = append(m.errorCalls, msg) }

// TestDivideLogsErrorOnZero — при делении на ноль логгер получает Error.
func TestDivideLogsErrorOnZero(t *testing.T) {
	mock := &mockLogger{}

	got, err := Divide(1, 0, mock)

	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Divide(1, 0) error = %v, want %v", err, ErrDivisionByZero)
	}
	if got != 0 {
		t.Errorf("Divide(1, 0) = %d, want 0", got)
	}
	if len(mock.errorCalls) != 1 {
		t.Fatalf("Error called %d times, want 1", len(mock.errorCalls))
	}
	if !strings.Contains(mock.errorCalls[0], "division by zero") {
		t.Errorf("Error message = %q, want contains %q", mock.errorCalls[0], "division by zero")
	}
	if len(mock.infoCalls) != 0 {
		t.Errorf("Info called %d times, want 0", len(mock.infoCalls))
	}
}

// TestDivideNoLogOnSuccess — при успешном делении Error не вызывается.
func TestDivideNoLogOnSuccess(t *testing.T) {
	mock := &mockLogger{}

	got, err := Divide(6, 3, mock)

	if err != nil {
		t.Fatalf("Divide(6, 3) unexpected error: %v", err)
	}
	if got != 2 {
		t.Errorf("Divide(6, 3) = %d, want 2", got)
	}
	if len(mock.errorCalls) != 0 {
		t.Errorf("Error called %d times, want 0", len(mock.errorCalls))
	}
}

// TestDivideNilLogger — nil-логгер допустим: логирование отключено.
func TestDivideNilLogger(t *testing.T) {
	got, err := Divide(1, 0, nil)

	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Divide(1, 0) error = %v, want %v", err, ErrDivisionByZero)
	}
	if got != 0 {
		t.Errorf("Divide(1, 0) = %d, want 0", got)
	}
}
