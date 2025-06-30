package items

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	contracts "item-detail-api/src/core/contracts/items"
	"item-detail-api/src/core/entities/items"
	usecase "item-detail-api/src/core/usecases/items"
	usecaseMocks "item-detail-api/src/core/usecases/items/mocks"
	"item-detail-api/src/entrypoints/rest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

const (
	listUrl = "/items"
)

var testItems = []items.Item{{
	ID:           uuid.New(),
	ProductID:    uuid.New(),
	ProviderID:   uuid.New(),
	SerialNumber: "SN123456",
	Stock:        10,
	CreatedAt:    time.Now(),
	UpdatedAt:    time.Now(),
	DeletedAt:    nil,
}}

func TestListItems_Handle(t *testing.T) {
	testCases := []rest.HandlerTestCase{
		{
			Name:       "Failed to bind",
			Request:    httptest.NewRequest(http.MethodGet, "/items?limit=asd", nil),
			Status:     http.StatusBadRequest,
			ResultBody: rest.Errf(CodeInvalidRequest, "Failed to bind query params"),
		},
		{
			Name:                   "Usecase ok",
			Request:                httptest.NewRequest(http.MethodGet, "/items?limit=10&offset=0", nil),
			Status:                 http.StatusOK,
			Usecase:                true,
			ResultBody:             rest.Okf(contracts.ToResponses(testItems, 1)),
			UsecaseArguments:       []interface{}{mock.Anything, usecase.ListItemsConfig{Limit: 10, Offset: 0}},
			UsecaseReturnArguments: []interface{}{testItems, 1, nil},
		},
		{
			Name:                   "Usecase Error",
			Request:                httptest.NewRequest(http.MethodGet, "/items?limit=10&offset=0", nil),
			Status:                 http.StatusInternalServerError,
			ResultBody:             rest.Errf(rest.CodeSomethingWentWrong, rest.DefaultErrorMessage),
			Usecase:                true,
			UsecaseArguments:       []interface{}{mock.Anything, usecase.ListItemsConfig{Limit: 10, Offset: 0}},
			UsecaseReturnArguments: []interface{}{[]items.Item{}, 0, errors.New("something went wrong")},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, runListItems(testCases[i]))
	}
}

func runListItems(tc rest.HandlerTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		mockUseCase := usecaseMocks.NewListItems(t)
		mockLogger := zaptest.NewLogger(t)
		handler := NewListItemsHandler(mockLogger, mockUseCase)
		handler.CustomNow = tc.UsecaseCustomNow
		router := rest.TestRouter(http.MethodGet, listUrl, handler.Handle(), uuid.New())

		if tc.Usecase {
			mockUseCase.
				On("Execute", tc.UsecaseArguments...).
				Return(tc.UsecaseReturnArguments...)
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, tc.Request)
		assert.Equal(t, tc.Status, w.Code)
		assert.Equal(t, tc.ResultBody.String(), w.Body.String())
	}
}

func TestNewWireListItemsHandler(t *testing.T) {
	mockLogger := zaptest.NewLogger(t)
	usecaseImpl := usecase.NewListItemsImpl(mockLogger, nil)
	handler := NewWireListItemsHandler(mockLogger, usecaseImpl)
	assert.NotNil(t, handler)
}
