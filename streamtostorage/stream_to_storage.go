package streamtostorage

import (
	"encoding/binary"
	"fmt"
	"io"
)

type MessageSizeBufferLen uint

const (
	MessageSizeBufferLen MessageSizeBufferLenSmall   = 2 // supports up to 65,536 bytes message size
	MessageSizeBufferLen MessageSizeBufferLenDefault = 4 // supports up to 4 GB message size
	MessageSizeBufferLen MessageSizeBufferLenLegacy  = 8 // supports up to 2 ^ 64 bytes message size (huge)
)

type putFuncType func(messageSize int) []byte

// Writer is a writer for writing streams of messages to a writer.
// The caller must provide synchronization for this writer; or use the SynchronizedWriter provided.
type Writer struct {
	file    io.Writer
	maxMessageSize int
	putFunc putFuncType
}

func NewWriter(file io.Writer, messageSizeBufferLen MessageSizeBufferLen) (*Writer, error) {
	putFunc, err := getPutFunc(messageSizeBufferLen)
	if err != nil {
		return nil, err
	}
	return &Writer{file,2 ^ (messageSizeBufferLen * 8), putFunc}
}

func getPutFunc(messageSizeBufferLen MessageSizeBufferLen) (putFuncType, error) {
	switch messageSizeBufferLen {
	case MessageSizeBufferLenSmall:
		return func(messageSize int) []byte {
			lenBuffer := make([]byte, messageSizeBufferLen)
			binary.LittleEndian.PutUint16(lenBuffer, uint16(messageSize))
			return lenBuffer
		}, nil

	case MessageSizeBufferLenDefault:
		return func(messageSize int) []byte {
			lenBuffer := make([]byte, messageSizeBufferLen)
			binary.LittleEndian.PutUint32(lenBuffer, uint32(messageSize))
			return lenBuffer
		}, nil

	case MessageSizeBufferLenLegacy:
		return func(messageSize int) []byte {
			lenBuffer := make([]byte, messageSizeBufferLen)
			binary.LittleEndian.PutUint64(lenBuffer, uint64(messageSize)
			return lenBuffer
		}, nil

	default:
		return nil, fmt.Errorf("unsupported size buffer len: %d", messageSizeBufferLen)
	}
}

func (s *Writer) Write(message []byte) (int, error) {
	messageLen :=len(message)

	if messageLen > s.maxMessageSize {
		return 0, fmt.Errorf("message larger than sizing allows; message is %d but max size is %d", messageLen, s.maxMessageSize)
	}

	lenBuffer := s.putFunc(messageLen)

	messageWithLen := append(lenBuffer, message...)

	return s.file.Write(messageWithLen)
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
	messageSizeBufferLen MessageSizeBufferLen
}

func NewReader(file io.Reader, messageSizeBufferLen MessageSizeBufferLen) *Reader {
	return &Reader{file, messageSizeBufferLen}
}

// ReadNextMessage reads the next message from the reader (starting from the beginning)
// Once the end of the reader has been reached, an io.EOF error is returned.
func (sr *Reader) ReadNextMessage() ([]byte, error) {
	lenBuffer := make([]byte, messageSizeBufferLen)
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
