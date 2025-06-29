package utils

import (
	"fmt"
	"net/http"
	"time"
)

type HandlerTestCase struct {
	Name                   string
	Request                *http.Request
	Status                 int
	ResultBody             fmt.Stringer
	ResultBodyBinary       []byte
	Usecase                bool
	UsecaseCustomNow       func() time.Time
	UsecaseArguments       []interface{}
	UsecaseReturnArguments []interface{}
}
