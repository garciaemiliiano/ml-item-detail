package ping

import (
	"item-detail-api/src/entrypoints/rest"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestPingHandler_Handle(t *testing.T) {
	testCases := []rest.HandlerTestCase{
		{
			Name:       "ok",
			Request:    httptest.NewRequest(http.MethodGet, "/", nil),
			Status:     http.StatusOK,
			ResultBody: rest.Okf("pong"),
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, run(testCases[i]))
	}
}

func run(tc rest.HandlerTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		mockLogger := zaptest.NewLogger(t)
		handler := NewPingHandler(mockLogger)
		router := rest.TestRouter(http.MethodGet, "/", handler.Handle(), uuid.New())

		w := httptest.NewRecorder()
		router.ServeHTTP(w, tc.Request)
		assert.Equal(t, tc.Status, w.Code)
		assert.Equal(t, tc.ResultBody.String(), w.Body.String())
	}
}

func TestNewPingHandler(t *testing.T) {
	mockLogger := zaptest.NewLogger(t)
	handler := NewPingHandler(mockLogger)
	assert.NotNil(t, handler)
}
