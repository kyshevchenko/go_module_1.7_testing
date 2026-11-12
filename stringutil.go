// Package stringutil предоставляет утилиты для работы со строками и числами.
package stringutil

import (
	"errors"
	"fmt"
	"sync"
)

// Reverse возвращает строку, записанную задом наперёд.
// Разворот выполняется по рунам, поэтому валидный UTF-8
// (включая кириллицу и эмодзи) обрабатывается корректно.
// Для невалидных UTF-8 байтов результат не определён:
// они заменяются на U+FFFD (см. тест FuzzReverse).
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ToUpper преобразует все символы строки в верхний регистр.
// Упрощение: корректно работает только для латиницы.
func ToUpper(s string) string {
	result := make([]rune, len([]rune(s)))
	for i, r := range []rune(s) {
		result[i] = r - 32
	}
	return string(result)
}

// Add складывает два целых числа.
func Add(a, b int) int {
	return a + b
}

// Divide делит a на b и возвращает ошибку при делении на ноль.
// Логгер опционален: nil отключает логирование.
func Divide(a, b int, log Logger) (int, error) {
	if b == 0 {
		if log != nil {
			log.Error(fmt.Sprintf("division by zero attempt: %d / %d", a, b))
		}
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// ErrDivisionByZero возвращается при попытке деления на ноль.
var ErrDivisionByZero = errors.New("division by zero")

// ConcurrentCounter инкрементирует счётчик из n горутин.
// Доступ к счётчику защищён мьютексом — гонки данных нет.
func ConcurrentCounter(n int) int {
	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	return counter
}
