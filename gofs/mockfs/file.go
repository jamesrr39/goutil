package mockfs

import (
	"bytes"
	"errors"
	"io"
)

type readSeeker struct {
	data     []byte
	position int64
}

func NewFile(data []byte) io.ReadSeeker {
	return readSeeker{data, 0}
}

func (rs readSeeker) Seek(position int64, whence int) (int64, error) {
	if whence != 0 {
		return 0, errors.New("unsupported whence method")
	}

	rs.position = position

	return position, nil
}

func (rs readSeeker) Read(byteSlice []byte) (int, error) {
	bb := bytes.NewBuffer(rs.data[rs.position:])
	return bb.Read(byteSlice)
}
