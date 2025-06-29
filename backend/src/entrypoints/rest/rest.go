package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestParam string
type EventRequestField string

const (
	Headers RequestParam = "headers"
	Uri     RequestParam = "uri"
	Body    RequestParam = "body"
	Form    RequestParam = "form"

	EventType    EventRequestField = "event_type"
	Payload      EventRequestField = "payload"
	ResourceID   EventRequestField = "resource_id"
	ResourceType EventRequestField = "resource_type"
)

var (
	errParamNotSpecified = errors.New("param not specified")
)

const CodeSomethingWentWrong = "F"
const DefaultErrorMessage = "Ocurrió un error"

func DefaultErrResponse() (status int, code string, message string) {
	return http.StatusInternalServerError, CodeSomethingWentWrong, DefaultErrorMessage
}

// ErrFormat is the default format used by all endpoints when returning an error.
type ErrFormat struct {
	Code string `json:"code"`
	Desc string `json:"description"`
}

func (f ErrFormat) String() string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}

func Errf(code, message string) ErrFormat {
	return ErrFormat{
		Code: code,
		Desc: message,
	}
}

// OkFormat is the default format used by all endpoints when returning an ok response.
type OkFormat struct {
	Content interface{} `json:"content"`
}

func (f OkFormat) String() string {
	bytes, _ := json.Marshal(f)
	return string(bytes)
}

func Okf(value ...interface{}) OkFormat {
	if len(value) == 0 {
		return OkFormat{}
	}
	return OkFormat{Content: value[0]}
}

// BindRequest binds the request parameters to the provided entity type T.
func BindRequest[T any](c *gin.Context, params ...RequestParam) (T, string, error) {
	var entity T
	if params == nil {
		return entity, "", errParamNotSpecified
	}

	for _, param := range params {
		switch param {
		case Headers:
			if err := c.ShouldBindHeader(&entity); err != nil {
				return entity, "Failed to bind headers", err
			}
		case Uri:
			if err := c.ShouldBindUri(&entity); err != nil {
				return entity, "Failed to bind uri", err
			}
		case Body:
			if err := c.ShouldBindJSON(&entity); err != nil {
				return entity, "Failed to bind body", err
			}
		case Form:
			if err := c.ShouldBindQuery(&entity); err != nil {
				return entity, "Failed to bind query params", err
			}
		}
	}

	return entity, "", nil
}
