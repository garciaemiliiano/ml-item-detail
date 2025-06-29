package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type HandlerTestCase struct {
	Name                   string
	Request                *http.Request
	Status                 int
	ResultBody             fmt.Stringer
	Usecase                bool
	UsecaseCustomNow       func() time.Time
	UsecaseArguments       []interface{}
	UsecaseReturnArguments []interface{}
}

func TestRouter(httpMethod, path string, handlerFunc gin.HandlerFunc, userID uuid.UUID, middleware ...gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Handle(httpMethod, path, handlerFunc)
	return router
}
