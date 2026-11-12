package stringutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// CalcResponseOf — вспомогательная функция: декодирует JSON-тело ответа.
func CalcResponseOf(t *testing.T, rec *httptest.ResponseRecorder) CalcResponse {
	t.Helper()
	var out CalcResponse
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&out))
	return out
}

// TestHandlerAddSuccess — успешное сложение через HTTP.
func TestHandlerAddSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc?op=add&a=2&b=3", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	out := CalcResponseOf(t, rec)
	assert.Equal(t, 5, out.Result)
	assert.Empty(t, out.Error)
}

// TestHandlerDivideSuccess — успешное деление через HTTP.
func TestHandlerDivideSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc?op=divide&a=6&b=3", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	out := CalcResponseOf(t, rec)
	assert.Equal(t, 2, out.Result)
}

// TestHandlerDivideByZero — деление на ноль: 400 и сообщение об ошибке.
func TestHandlerDivideByZero(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc?op=divide&a=1&b=0", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	out := CalcResponseOf(t, rec)
	assert.Contains(t, out.Error, "division by zero")
}

// TestHandlerUnknownOp — неизвестная операция: 400.
func TestHandlerUnknownOp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc?op=multiply&a=2&b=3", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	out := CalcResponseOf(t, rec)
	assert.Contains(t, out.Error, "unsupported op")
}

// TestHandlerInvalidParams — нечисловые параметры: 400.
func TestHandlerInvalidParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/calc?op=add&a=abc&b=3", nil)
	rec := httptest.NewRecorder()

	Handler(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	out := CalcResponseOf(t, rec)
	assert.Contains(t, out.Error, "invalid a or b")
}

// TestClientDo — клиент против httptest.NewServer: настоящий HTTP round-trip.
func TestClientDo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Handler))
	defer server.Close()

	client := NewClient(server.URL)

	got, err := client.Do("add", 2, 3)
	require.NoError(t, err)
	assert.Equal(t, 5, got)
}

// TestClientDoDivideError — клиент прокидывает ошибку сервера.
func TestClientDoDivideError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Handler))
	defer server.Close()

	client := NewClient(server.URL)

	_, err := client.Do("divide", 1, 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "division by zero")
}
