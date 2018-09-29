package streamtostorage

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	message1 = "hello 1"
	message2 = "hello message 2 with different length"
)

func Test_WriteRead(t *testing.T) {
	bb := bytes.NewBuffer(nil)
	writer := NewStreamToStorageWriter(bb)
	err := writer.Write([]byte(message1))
	require.Nil(t, err)
	err = writer.Write([]byte(message2))
	require.Nil(t, err)

	reader := NewStreamToStorageReader(bb)
	message, err := reader.ReadNextMessage()
	require.Nil(t, err)
	assert.Equal(t, []byte(message1), message)

	message, err = reader.ReadNextMessage()
	require.Nil(t, err)
	assert.Equal(t, []byte(message2), message)

	_, err = reader.ReadNextMessage()
	assert.Equal(t, io.EOF, err)
}
