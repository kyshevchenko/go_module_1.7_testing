package stringutil

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// calcResponse — формат JSON-ответа калькулятора.
type calcResponse struct {
	Result int    `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Handler — HTTP-обработчик калькулятора.
// Параметры запроса: op (add | divide), a, b — целые числа.
// Ответ: JSON {"result": N} при успехе или {"error": "..."} при ошибке.
func Handler(w http.ResponseWriter, r *http.Request) {
	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))
	if errA != nil || errB != nil {
		writeCalcError(w, http.StatusBadRequest, "invalid a or b: expected integers")
		return
	}

	op := r.URL.Query().Get("op")

	var result int
	switch op {
	case "add":
		result = Add(a, b)
	case "divide":
		var err error
		result, err = Divide(a, b, SimpleLogger{})
		if err != nil {
			writeCalcError(w, http.StatusBadRequest, err.Error())
			return
		}
	default:
		writeCalcError(w, http.StatusBadRequest, fmt.Sprintf("unsupported op %q: use add or divide", op))
		return
	}

	writeCalcJSON(w, http.StatusOK, calcResponse{Result: result})
}

// writeCalcJSON — отправка успешного JSON-ответа.
func writeCalcJSON(w http.ResponseWriter, status int, resp calcResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

// writeCalcError — отправка JSON-ответа с ошибкой.
func writeCalcError(w http.ResponseWriter, status int, msg string) {
	writeCalcJSON(w, status, calcResponse{Error: msg})
}

// Client — HTTP-клиент калькулятора.
type Client struct {
	BaseURL string
	HTTP    *http.Client
}

// NewClient создаёт клиент с http-клиентом по умолчанию.
func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL, HTTP: http.DefaultClient}
}

// Do выполняет операцию op над a и b через HTTP API калькулятора.
func (c *Client) Do(op string, a, b int) (int, error) {
	url := fmt.Sprintf("%s/calc?op=%s&a=%d&b=%d", c.BaseURL, op, a, b)

	resp, err := c.HTTP.Get(url)
	if err != nil {
		return 0, fmt.Errorf("calc request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	var out calcResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, fmt.Errorf("decode response: %w", err)
	}
	if out.Error != "" {
		return 0, fmt.Errorf("calc: %s", out.Error)
	}
	return out.Result, nil
}
