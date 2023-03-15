package response

import (
	"encoding/json"
	"net/http"
)

type SuccessContent struct {
	httpStatusCode int
	data           any
}

func Success(httpStatusCode int, data any) SuccessContent {
	return SuccessContent{
		httpStatusCode: httpStatusCode,
		data:           data,
	}
}

func Created(data any) SuccessContent {
	return Success(http.StatusCreated, data)
}

func NoContent() SuccessContent {
	return Success(http.StatusNoContent, nil)
}

func OK(data any) SuccessContent {
	return Success(http.StatusOK, data)
}

func (c SuccessContent) Kind() string {
	return "success"
}

func (c SuccessContent) HttpStatusCode() int {
	return c.httpStatusCode
}

func (c SuccessContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.data)
}
