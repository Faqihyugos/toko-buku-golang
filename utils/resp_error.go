package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Status  int
	Message string
	Data    interface{}
}

type MessageErr interface {
	Message() string
	Status() int
	Error() string
}

type messageErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *messageErr) Error() string {
	return e.ErrError
}

func (e *messageErr) Message() string {
	return e.ErrMessage
}

func (e *messageErr) Status() int {
	return e.ErrStatus
}
func NewNotFoundError(message string) MessageErr {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewBadRequestError(message string) MessageErr {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrError:   "invalid_request",
	}
}

func NewApiErrFromBytes(body []byte) (MessageErr, error) {
	var result messageErr
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func NewInternalServerError(message string) MessageErr {
	return &messageErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "server_error",
	}
}

func Response(code int, message string, data interface{}) ResponseMessage {
	return ResponseMessage{Status: code, Message: message, Data: data}
}