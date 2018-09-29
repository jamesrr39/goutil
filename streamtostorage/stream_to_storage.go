package streamtostorage

import (
	"encoding/binary"
	"io"
)

// StreamToStorageWriter is a writer for writing streams of messages to a writer.
// Writes are not synchronized, the caller must provide the synchronization if it will be written to from multiple goroutines.
type StreamToStorageWriter struct {
	file io.Writer
}

func NewStreamToStorageWriter(file io.Writer) *StreamToStorageWriter {
	return &StreamToStorageWriter{file}
}

func (s *StreamToStorageWriter) Write(message []byte) error {
	lenBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(lenBuffer, uint64(len(message)))

	_, err := s.file.Write(append(lenBuffer, message...))
	return err
}

type StreamToStorageReader struct {
	file io.Reader
}

func NewStreamToStorageReader(file io.Reader) *StreamToStorageReader {
	return &StreamToStorageReader{file}
}

// ReadNextMessage reads the next message from the reader (starting from the beginning)
// Once the end of the reader has been reached, an io.EOF error is returned.
func (sr *StreamToStorageReader) ReadNextMessage() ([]byte, error) {
	lenBuffer := make([]byte, 8)
	_, err := sr.file.Read(lenBuffer)
	if err != nil {
		return nil, err
	}

	messageLen := binary.LittleEndian.Uint64(lenBuffer)
	messageBuffer := make([]byte, messageLen)

	_, err = sr.file.Read(messageBuffer)
	if err != nil {
		return nil, err
	}

	return messageBuffer, nil
}
