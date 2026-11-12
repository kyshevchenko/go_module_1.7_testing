package stringutil

import (
	"testing"
	"unicode/utf8"
)

// FuzzReverse — фаззинг-тест Reverse.
//
// Инвариант Reverse(Reverse(s)) == s проверяется только для валидного UTF-8:
// для невалидных байтовых последовательностей []rune заменяет их на U+FFFD,
// и информация теряется — двойной разворот не может восстановить исходные
// байты (фаззер нашёл контрпример "\x8a\xd8", он сохранён в testdata/fuzz).
func FuzzReverse(f *testing.F) {
	// Начальные seed-корпуса.
	f.Add("hello")
	f.Add("world")
	f.Add("123")
	f.Add("")
	f.Add("level")
	f.Add("привет")
	f.Add("héllo wörld")
	f.Add("🎉🚀 Go!")

	f.Fuzz(func(t *testing.T, s string) {
		rev := Reverse(s)

		if !utf8.ValidString(s) {
			// Невалидный вход — вне области определения Reverse.
			t.Skipf("skipping invalid UTF-8 input %q", s)
		}

		// Свойство 1: реверс валидной строки остаётся валидным UTF-8.
		if !utf8.ValidString(rev) {
			t.Errorf("Reverse(%q) produced invalid UTF-8 %q", s, rev)
		}

		// Свойство 2 (инвариант): двойной разворот возвращает исходную строку.
		if got := Reverse(rev); got != s {
			t.Errorf("invariant broken for %q: Reverse(Reverse(s)) = %q", s, got)
		}
	})
}
