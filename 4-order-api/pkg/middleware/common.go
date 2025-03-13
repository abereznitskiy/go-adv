package middleware

import "net/http"

type WrapperWriter struct {
	http.ResponseWriter
	StatusCode   int
	ResponseSize int
}

func (w *WrapperWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func (w *WrapperWriter) Write(b []byte) (int, error) {
	num, err := w.ResponseWriter.Write(b)
	w.ResponseSize += num
	return num, err
}
