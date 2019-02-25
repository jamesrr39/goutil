package errorsx

import (
	"net/http"
)

type logger interface {
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

func HTTPError(w http.ResponseWriter, log logger, err *Error, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode < 500 {
		log.Warn(err.Error())
	} else {
		log.Error(err.Error())
	}

	w.Write([]byte(err.message()))
}
