package response

import (
	"encoding/json"
	"net/http"
)

type ErrorContent struct {
	httpStatusCode int
	message        string
}

func Error(httpStatusCode int, err error) ErrorContent {
	return ErrorContent{
		httpStatusCode: httpStatusCode,
		message:        err.Error(),
	}
}

func BadRequest(err error) ErrorContent {
	return Error(http.StatusBadRequest, err)
}

func NotFound(err error) ErrorContent {
	return Error(http.StatusNotFound, err)
}

func InternalServerError(err error) ErrorContent {
	return Error(http.StatusInternalServerError, err)
}

func Unauthorized(err error) ErrorContent {
	return Error(http.StatusUnauthorized, err)
}

func (c ErrorContent) Kind() string {
	return "error"
}

func (c ErrorContent) HttpStatusCode() int {
	return c.httpStatusCode
}

func (c ErrorContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"message": c.message,
	})
}
