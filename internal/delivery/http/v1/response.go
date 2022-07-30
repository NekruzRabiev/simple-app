package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nekruzrabiev/simple-app/pkg/logger"
	"net/http"
)

type errResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}

func newErrResponse(message string, code int) *errResponse {
	return &errResponse{
		Message: message,
		Code:    code,
	}
}

func (e *errResponse) marshalJSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}

var (
	ErrBadParams                  = newErrResponse("некорректные параметры запроса", http.StatusBadRequest)
	ErrSession                    = newErrResponse("вы не авторизованы, пожалуйста перезайдите", http.StatusUnauthorized)
	ErrInternalServer             = newErrResponse("упс, что-то пошло не так", http.StatusInternalServerError)
	ErrUserExists                 = newErrResponse("пользователь уже существует", http.StatusConflict)
	ErrBadEmailOrPassword         = newErrResponse("некорректный логин или пароль", http.StatusForbidden)
	ErrNotContainsDigitAndLetters = newErrResponse("пароль должен содержать цифры и буквы", http.StatusBadRequest)
)

func newResponse(c *gin.Context, statusCode int, resp *errResponse, err error) {
	message := err.Error()
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, resp)
}
