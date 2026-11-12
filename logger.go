package stringutil

import (
	"fmt"
	"os"
)

// Logger — минимальный интерфейс логирования.
// Позволяет подставлять реализации (реальный, мок) в тестируемый код.
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// SimpleLogger — реализация Logger, пишущая в os.Stdout.
type SimpleLogger struct{}

// Info пишет информационное сообщение в stdout.
func (SimpleLogger) Info(msg string) {
	fmt.Fprintln(os.Stdout, "INFO:", msg)
}

// Error пишет сообщение об ошибке в stdout.
func (SimpleLogger) Error(msg string) {
	fmt.Fprintln(os.Stdout, "ERROR:", msg)
}
