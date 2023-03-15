package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, content Content) {
	if b, err := json.Marshal(content); err != nil {
		JSON(w, InternalServerError(err))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(content.HttpStatusCode())
		w.Write(b)
	}
}

type Content interface {
	Kind() string
	HttpStatusCode() int
}
