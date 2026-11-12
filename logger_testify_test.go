package stringutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// testifyMockLogger — мок Logger на базе testify/mock.
// Каждый вызов метода пробрасывается в mock.Mock, который
// фиксирует факт вызова и переданные аргументы.
type testifyMockLogger struct {
	mock.Mock
}

func (m *testifyMockLogger) Info(msg string)  { m.Called(msg) }
func (m *testifyMockLogger) Error(msg string) { m.Called(msg) }

// TestDivideZeroLogsErrorTestify — testify-мок: Error вызван с точным сообщением.
func TestDivideZeroLogsErrorTestify(t *testing.T) {
	log := new(testifyMockLogger)
	log.On("Error", "division by zero attempt: 1 / 0").Return().Once()

	got, err := Divide(1, 0, log)

	require.ErrorIs(t, err, ErrDivisionByZero) // дальше нет смысла, если ошибка не та
	assert.Equal(t, 0, got)
	log.AssertExpectations(t) // все .Once() вызовы произошли, лишних нет
}

// TestDivideSuccessNoErrorLogTestify — при успехе Error не вызывается.
func TestDivideSuccessNoErrorLogTestify(t *testing.T) {
	log := new(testifyMockLogger)

	got, err := Divide(6, 3, log)

	require.NoError(t, err)
	assert.Equal(t, 2, got)
	log.AssertNotCalled(t, "Error", mock.Anything)
	log.AssertNotCalled(t, "Info", mock.Anything)
}
