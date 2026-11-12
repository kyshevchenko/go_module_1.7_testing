package stringutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSaveToFile — запись в файл во временной директории и проверка содержимого.
func TestSaveToFile(t *testing.T) {
	dir := t.TempDir() // создаётся до теста, автоматически удаляется после
	path := filepath.Join(dir, "result.txt")

	data := []byte("hello, file!")
	require.NoError(t, SaveToFile(path, data))

	// Файл существует.
	info, err := os.Stat(path)
	require.NoError(t, err, "file should exist after SaveToFile")
	assert.Equal(t, int64(len(data)), info.Size())

	// Содержимое совпадает.
	got, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, data, got)
}

// TestSaveToFileOverwrite — повторная запись перезаписывает содержимое.
func TestSaveToFileOverwrite(t *testing.T) {
	path := filepath.Join(t.TempDir(), "over.txt")

	require.NoError(t, SaveToFile(path, []byte("first")))
	require.NoError(t, SaveToFile(path, []byte("second")))

	got, err := os.ReadFile(path)
	require.NoError(t, err)
	assert.Equal(t, []byte("second"), got)
}

// TestSaveToFileBadPath — запись в несуществующую директорию возвращает ошибку.
func TestSaveToFileBadPath(t *testing.T) {
	path := filepath.Join(t.TempDir(), "no", "such", "dir", "file.txt")

	err := SaveToFile(path, []byte("data"))
	require.Error(t, err)
}
