package middleware

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

const (
	AppName = `graphblog`
)

type Middleware struct {
	log *zap.Logger
}

func NewMiddleware(logger *zap.Logger) IMiddleware {
	return &Middleware{
		log: logger,
	}
}

func (mw *Middleware) Wrap(h http.Handler) http.Handler {
	return mw.requestLogger(h)
}

func (mw *Middleware) requestLogger(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now().UnixMilli()

		capture := httptest.NewRecorder()

		h.ServeHTTP(capture, r)

		CopyResponseHeadersToResponseWriter(capture.Result(), w)
		written, err := io.Copy(w, capture.Result().Body)
		if err != nil {
			mw.log.Error(fmt.Sprintf("%d bytes written - error: %v", written, err))
		}

		endTime := time.Now().UnixMilli()
		requestTime := endTime - startTime

		mw.log.Info(`trace`,
			zap.String(`app`, AppName),
			zap.String(`path`, r.URL.Path),
			zap.Int64(`rt`, requestTime),
			zap.Int(`status`, capture.Result().StatusCode),
			zap.String(`content-type`, capture.Result().Header.Get(`Content-Type`)),
		)
	})

}
