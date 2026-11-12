package stringutil

import "os"

// SaveToFile записывает данные в файл по указанному пути.
// Файл создаётся с правами 0644; существующий файл перезаписывается.
func SaveToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0o644)
}
