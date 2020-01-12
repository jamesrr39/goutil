package httpextra

import "net/http"

// Doer is implemented by http.DefaultClient
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type MockDoer struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (d *MockDoer) Do(req *http.Request) (*http.Response, error) {
	return d.DoFunc(req)
}
