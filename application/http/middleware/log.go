package middleware

import (
	"net/http"
	"time"

	"convercy/domain/valueobject"
	"convercy/shared/logging"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func Log(logger logging.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				now       = time.Now()
				requestID = valueobject.GenerateRequestID()
			)

			logger.Info("Got request",
				logging.String("http_method", r.Method),
				logging.String("request_uri", r.RequestURI),
				logging.Stringer("request_id", requestID),
			)

			rw := negroni.NewResponseWriter(w)

			next.ServeHTTP(rw, r)

			logger.Info("Returning response",
				logging.Stringer("request_id", requestID),
				logging.Int("http_status_code", rw.Status()),
				logging.Stringer("response_time", time.Since(now)),
			)
		})
	}
}
