package items

import (
	"errors"
	"fmt"
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

var (
	url      = "/items/:id"
	itemID   = uuid.New()
	testItem = items.Item{
		ID:           itemID,
		ProductID:    uuid.New(),
		ProviderID:   uuid.New(),
		SerialNumber: "SN123456",
		Stock:        10,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}
)

func TestGetItemHandler_Handle(t *testing.T) {
	testCases := []rest.HandlerTestCase{
		{
			Name:       "Failed to bind",
			Request:    httptest.NewRequest(http.MethodGet, "/items/1", nil),
			Status:     http.StatusBadRequest,
			ResultBody: rest.Errf(CodeInvalidRequest, "Failed to bind uri"),
		},
		{
			Name:                   "Usecase ok",
			Request:                httptest.NewRequest(http.MethodGet, fmt.Sprintf("/items/%s", itemID), nil),
			Status:                 http.StatusOK,
			Usecase:                true,
			ResultBody:             rest.Okf(contracts.ToResponse(testItem)),
			UsecaseArguments:       []interface{}{mock.Anything, usecase.GetItemParams{ID: itemID}},
			UsecaseReturnArguments: []interface{}{testItem, nil},
		},
		{
			Name:                   "Usecase Error",
			Request:                httptest.NewRequest(http.MethodGet, fmt.Sprintf("/items/%s", itemID), nil),
			Status:                 http.StatusInternalServerError,
			ResultBody:             rest.Errf(rest.CodeSomethingWentWrong, rest.DefaultErrorMessage),
			Usecase:                true,
			UsecaseArguments:       []interface{}{mock.Anything, usecase.GetItemParams{ID: itemID}},
			UsecaseReturnArguments: []interface{}{items.Item{}, errors.New("something went wrong")},
		},
		{
			Name:                   "Usecase Error Item Required",
			Request:                httptest.NewRequest(http.MethodGet, fmt.Sprintf("/items/%s", uuid.Nil), nil),
			Status:                 http.StatusBadRequest,
			ResultBody:             rest.Errf("err-item-id-required", "item ID is required"),
			Usecase:                true,
			UsecaseArguments:       []interface{}{mock.Anything, usecase.GetItemParams{ID: uuid.Nil}},
			UsecaseReturnArguments: []interface{}{items.Item{}, usecase.ErrItemIDRequired},
		},
		{
			Name:                   "Usecase Error item not found",
			Request:                httptest.NewRequest(http.MethodGet, fmt.Sprintf("/items/%s", itemID), nil),
			Status:                 http.StatusBadRequest,
			ResultBody:             rest.Errf("err-item-not-found", "item not found"),
			Usecase:                true,
			UsecaseArguments:       []interface{}{mock.Anything, usecase.GetItemParams{ID: itemID}},
			UsecaseReturnArguments: []interface{}{items.Item{}, usecase.ErrItemNotFound},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, run(testCases[i]))
	}
}

func run(tc rest.HandlerTestCase) func(t *testing.T) {
	return func(t *testing.T) {
		mockUseCase := usecaseMocks.NewGetItem(t)
		mockLogger := zaptest.NewLogger(t)
		handler := NewGetItemHandler(mockLogger, mockUseCase)
		handler.CustomNow = tc.UsecaseCustomNow
		router := rest.TestRouter(http.MethodGet, url, handler.Handle(), uuid.New())

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

func TestNewWireGetItemHandler(t *testing.T) {
	mockLogger := zaptest.NewLogger(t)
	usecaseImpl := usecase.NewGetItemImpl(mockLogger, nil)
	handler := NewWireGetItemHandler(mockLogger, usecaseImpl)
	assert.NotNil(t, handler)
}
