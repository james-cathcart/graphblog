package middleware

import "net/http"

func CopyResponseHeadersToResponseWriter(source *http.Response, dest http.ResponseWriter) {
	for header, values := range source.Header {
		for _, value := range values {
			dest.Header().Add(header, value)
		}
	}
}
