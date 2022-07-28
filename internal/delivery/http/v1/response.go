package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nekruzrabiev/simple-app/pkg/logger"
	"net/http"
)

type dataResponse struct {
	Data interface{} `json:"data"`
}

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

var (
	ErrBadParams                  = newErrResponse("некорректные параметры запроса", http.StatusBadRequest)
	ErrSession                    = newErrResponse("вы не авторизованы, пожалуйста перезайдите", http.StatusUnauthorized)
	ErrInternalServer             = newErrResponse("упс, что-то пошло не так", http.StatusInternalServerError)
	ErrInsertLimit                = newErrResponse("Превышен лимит адресов", http.StatusMethodNotAllowed)
	ErrDeniedCancelOrder          = newErrResponse("Нельзя отменить заказ", http.StatusMethodNotAllowed)
	ErrPriceMismatch              = newErrResponse("Цены усторели, пожалуйста, очистите корзину и сделайте заново заказ", http.StatusConflict)
	ErrItemPriceMismatch          = newErrResponse("цена продукта не соответствует", http.StatusConflict)
	ErrInvalidPhoneNumber         = newErrResponse("Неверный ввод номера телефона", http.StatusForbidden)
	ErrNoFcmTopic                 = newErrResponse("некорректный топик", http.StatusBadRequest)
	ErrSubscribeUnsubscribe       = newErrResponse("не получилось изменить статус нотификаций", http.StatusConflict)
	ErrTurnOffNotifications       = newErrResponse("пользователь выключил уведомления", http.StatusConflict)
	ErrUserExists                 = newErrResponse("пользователь уже существует", http.StatusConflict)
	ErrBadPhoneOrPassword         = newErrResponse("некорректный логин или пароль", http.StatusForbidden)
	ErrRestDate                   = newErrResponse("выберите дату минимум за два дня", http.StatusMethodNotAllowed)
	ErrUnfinishedOrder            = newErrResponse("У Вас незавершенный заказ", http.StatusForbidden)
	ErrOldPassword                = newErrResponse("некорректный старый пароль", http.StatusForbidden)
	ErrNotContainsDigitAndLetters = newErrResponse("пароль должен содержать цифры и буквы", http.StatusForbidden)
	ErrOrderStatus                = newErrResponse("некорректный статус", http.StatusBadRequest)
	ErrCourierSession             = newErrResponse("курьер в оффлайне", http.StatusConflict)
	ErrUserNotExist               = newErrResponse("пользователь не существует", http.StatusNotFound)
	ErrInputFileFormat            = newErrResponse("неверное расширение документа, только pdf", http.StatusConflict)
	ErrOneItem                    = newErrResponse("Нельзя удалить последний продукт в заказе", http.StatusConflict)
	ErrOTPCodeNotMatch            = newErrResponse("Неверный код активации", http.StatusForbidden)
	ErrNotify                     = newErrResponse("не получилось отправить уведомление, попробуйте еще", 444)
)

func newResponse(c *gin.Context, statusCode int, resp *errResponse, err error) {
	message := err.Error()
	logger.Error(message)
	c.AbortWithStatusJSON(statusCode, resp)
}