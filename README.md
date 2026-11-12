# stringutil

[![Go Version](https://img.shields.io/badge/go-1.26-blue.svg)](https://go.dev/)
[![Go Reference](https://pkg.go.dev/badge.svg)](https://pkg.go.dev/)
[![Lint](https://img.shields.io/badge/lint-golangci--lint-brightgreen.svg)](.golangci.yml)
[![Coverage](https://img.shields.io/badge/coverage-95%25-success.svg)](#тесты)

Учебный пакет Go из курса [go_module_1.7_testing](https://github.com/kyshevchenko/go_module_1.7_testing):
утилиты для работы со строками и числами, от которых требуется корректное
поведение под тестами, бенчмарками и фаззингом.

## Возможности

| Функция | Описание |
|---|---|
| `Reverse(s string) string` | разворот строки по рунам (unicode-safe) |
| `ToUpper(s string) string` | верхний регистр (учебное упрощение — только латиница) |
| `Add(a, b int) int` | сложение целых чисел |
| `Divide(a, b int, log Logger) (int, error)` | целочисленное деление с ошибкой `ErrDivisionByZero` и опциональным логгером |
| `ConcurrentCounter(n int) int` | счётчик из n горутин, защищённый мьютексом |
| `SaveToFile(filename string, data []byte) error` | запись данных в файл |
| `Handler` / `Client.Do` | HTTP-калькулятор (`op=add|divide`) и клиент к нему |

## Установка

Требуется Go 1.18+ (фаззинг), фактически разрабатывалось на Go 1.26.

```bash
git clone https://github.com/kyshevchenko/go_module_1.7_testing.git
cd go_module_1.7_testing/7.test.stringutil
go build ./...
```

## Примеры использования

```go
package main

import (
	"fmt"

	"stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("hello")) // olleh

	result, err := stringutil.Divide(6, 3, nil)
	fmt.Println(result, err) // 2 <nil>

	_, err = stringutil.Divide(1, 0, nil)
	fmt.Println(err) // division by zero
}
```

## Тесты

```bash
make test    # все тесты с -race
make cover   # покрытие + HTML-отчёт в браузере
make lint    # golangci-lint
make fmt     # gofmt + goimports
make clean   # удалить временные файлы
```

Без Makefile:

```bash
go test -race ./...              # тесты (юнит, табличные, race, fuzz-регресс)
go test -bench=. -benchmem ./... # бенчмарки
go test -fuzz=FuzzReverse -fuzztime=30s .  # фаззинг
```

## Документация

```bash
go doc stringutil.Divide      # в терминале
godoc -http=:6060             # в браузере: http://localhost:6060/pkg/stringutil/
```

Godoc-комментарии + исполняемые примеры (`ExampleReverse`, `ExampleDivide`) —
в файлах пакета; примеры проверяются `go test`.

## Структура проекта

```
7.test.stringutil/
├── stringutil.go        # основные функции: Reverse, ToUpper, Add, Divide, ConcurrentCounter
├── logger.go            # интерфейс Logger и SimpleLogger
├── handler.go           # HTTP-калькулятор и клиент (swag-аннотации)
├── fileio.go            # SaveToFile
├── cmd/calcserver/      # демо-сервер: /calc + Swagger UI на :8081
├── docs/                # сгенерированная swagger-документация (swag init)
├── *_test.go            # юнит, табличные, race, benchmark, fuzz, example-тесты
├── testdata/fuzz/       # регресс-корпус фаззера (найденные кейсы)
├── .golangci.yml        # конфигурация линтера
└── Makefile             # build / test / cover / lint / fmt / clean
```

## Swagger

```bash
go run ./cmd/calcserver
# API:   http://localhost:8081/calc?op=add&a=2&b=3
# UI:    http://localhost:8081/swagger/index.html
```

Регенерация документации после изменения аннотаций:
`~/go/bin/swag init -g cmd/calcserver/main.go -o docs`
