# Makefile для пакета stringutil (Задание 6).

GOBIN := $(HOME)/go/bin

.PHONY: build test cover lint fmt clean

# В пакете нет main — build просто проверяет компиляцию всех пакетов.
build:
	go build ./... && echo "✓ Сборка прошла успешно"

test:
	go test -race ./... && echo "✓ Все тесты прошли"

cover:
	go test -coverprofile=c.out ./...
	go tool cover -func=c.out | tail -1
	go tool cover -html=c.out

lint:
	$(GOBIN)/golangci-lint run && echo "✓ Линтер не нашёл ошибок"

fmt:
	gofmt -w .
	$(GOBIN)/goimports -w .

clean:
	rm -f c.out coverage.html *.test
	echo "✓ Временные файлы удалены"
