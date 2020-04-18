package ioextra

import "io"

type CustomCloser struct {
	io.Reader
	closeFunc func() error
}

func NewCustomCloser(reader io.Reader, closeFunc func() error) io.ReadCloser {
	return CustomCloser{reader, closeFunc}
}

func (cc CustomCloser) Close() error {
	return cc.closeFunc()
}
