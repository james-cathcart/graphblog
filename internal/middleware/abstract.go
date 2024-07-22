package middleware

import "net/http"

type IMiddleware interface {
	Wrap(handler http.Handler) http.Handler
}
