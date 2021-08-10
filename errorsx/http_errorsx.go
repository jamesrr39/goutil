package errorsx

import (
	"encoding/json"
	"net/http"
)

type logger interface {
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

func HTTPError(w http.ResponseWriter, log logger, err Error, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode < 500 {
		log.Warn("%s. Stack trace:\n%s", err.Error(), err.Stack())
	} else {
		log.Error("%s. Stack trace:\n%s", err.Error(), err.Stack())
	}

	w.Write([]byte(err.Error()))
}

type jsonErrorMessageType struct {
	Message string `json:"message"`
}

func HTTPJSONError(w http.ResponseWriter, log logger, err Error, statusCode int) {
	w.WriteHeader(statusCode)
	if statusCode < 500 {
		log.Warn("%s. Stack trace:\n%s", err.Error(), err.Stack())
	} else {
		log.Error("%s. Stack trace:\n%s", err.Error(), err.Stack())
	}

	json.NewEncoder(w).Encode(jsonErrorMessageType{err.Error()})
}
