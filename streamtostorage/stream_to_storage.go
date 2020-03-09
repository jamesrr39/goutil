package streamtostorage

import (
	"encoding/binary"
	"io"
)

// Writer is a writer for writing streams of messages to a writer.
// The caller must provide synchronization for this writer; or use the SynchronizedWriter provided.
type Writer struct {
	file io.Writer
}

func NewWriter(file io.Writer) *Writer {
	return &Writer{file}
}

func (s *Writer) Write(message []byte) (int, error) {
	messageWithLen := makeMessageWithLen(message)

	return s.file.Write(messageWithLen)
}

func makeMessageWithLen(message []byte) []byte {

	lenBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(lenBuffer, uint64(len(message)))

	messageWithLen := append(lenBuffer, message...)

	return messageWithLen
}

type synchronizedWriteMessage struct {
	MessageWithLen []byte
	OnFinishedChan chan (respType)
}

type SynchronizedWriter struct {
	writer    io.Writer
	writeChan chan (*synchronizedWriteMessage)
}

func NewSynchronizedWriter(file io.Writer) *SynchronizedWriter {
	w := &SynchronizedWriter{file, make(chan (*synchronizedWriteMessage))}

	go func() {
		for {
			message := <-w.writeChan
			c, err := w.writer.Write(message.MessageWithLen)

			message.OnFinishedChan <- respType{c, err}
		}
	}()

	return w
}

type respType struct {
	count int
	err   error
}

func (s *SynchronizedWriter) Write(message []byte) (int, error) {
	syncMessage := &synchronizedWriteMessage{
		MessageWithLen: makeMessageWithLen(message),
		OnFinishedChan: make(chan (respType)),
	}

	s.writeChan <- syncMessage

	resp := <-syncMessage.OnFinishedChan

	return resp.count, resp.err
}

type Reader struct {
	file io.Reader
}

func NewReader(file io.Reader) *Reader {
	return &Reader{file}
}

// ReadNextMessage reads the next message from the reader (starting from the beginning)
// Once the end of the reader has been reached, an io.EOF error is returned.
func (sr *Reader) ReadNextMessage() ([]byte, error) {
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
