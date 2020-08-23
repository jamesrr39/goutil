package streamtostorage

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	message1        = "hello 1"
	message2        = "hello message 2 with different length"
	uint64SizeBytes = 8
)

func Test_WriteRead(t *testing.T) {
	bb := bytes.NewBuffer(nil)
	writer, err := NewWriter(bb, MessageSizeBufferLenDefault)
	require.Nil(t, err)

	// write messages
	i, err := writer.Write([]byte(message1))
	require.Nil(t, err)
	require.Equal(t, len(message1), i)

	i, err = writer.Write([]byte(message2))
	require.Nil(t, err)
	require.Equal(t, len(message2), i)

	// read messages
	reader := NewReader(bb, MessageSizeBufferLenDefault)
	message, err := reader.ReadNextMessage()
	require.Nil(t, err)
	assert.Equal(t, []byte(message1), message)

	message, err = reader.ReadNextMessage()
	require.Nil(t, err)
	assert.Equal(t, []byte(message2), message)

	_, err = reader.ReadNextMessage()
	assert.Equal(t, io.EOF, err)
}

func TestSynchronizedWriter_Write(t *testing.T) {
	sw, err := NewSynchronizedWriter(bytes.NewBuffer(nil), MessageSizeBufferLenDefault)
	require.NoError(t, err)

	type args struct {
		message []byte
	}
	tests := []struct {
		name    string
		writer  *SynchronizedWriter
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "simple write",
			writer: sw,
			args: args{
				message: []byte("hello"),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.writer.Write(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SynchronizedWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

// Test_synchronizedWrites tests concurrent synchronized writes
func Test_synchronizedWrites(t *testing.T) {
	bb := bytes.NewBuffer(nil)

	w, err := NewSynchronizedWriter(bb, MessageSizeBufferLenDefault)
	require.NoError(t, err)

	haveSeenNumYetMap := make(map[string]bool)

	bottomLimit := 0
	topLimit := 50

	var wg sync.WaitGroup
	for i := bottomLimit; i < topLimit; i++ {
		wg.Add(1)

		val := fmt.Sprintf("%d", i)

		haveSeenNumYetMap[val] = false

		go func() {
			defer wg.Done()
			_, err := w.Write([]byte(val))
			require.NoError(t, err)
		}()
	}

	wg.Wait()

	r := NewReader(bb, MessageSizeBufferLenDefault)
	for {
		b, err := r.ReadNextMessage()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)

		val, err := strconv.Atoi(string(b))
		require.NoError(t, err)

		assert.True(t, val >= bottomLimit)
		assert.True(t, val < topLimit)
		assert.Equal(t, false, haveSeenNumYetMap[string(b)])

		// all checks done, mark value as seen
		haveSeenNumYetMap[string(b)] = true
	}

	for k, v := range haveSeenNumYetMap {
		assert.True(t, v, "hasn't seen val %v", k)
	}
}

func Test_getMaxMessageSize(t *testing.T) {
	type args struct {
		messageSizeBufferLen MessageSizeBufferLen
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "MessageSizeBufferLenSmall",
			args: args{MessageSizeBufferLenSmall},
			want: 65536,
		}, {
			name: "MessageSizeBufferLenDefault",
			args: args{MessageSizeBufferLenDefault},
			want: 4294967296,
		}, {
			name: "MessageSizeBufferLenLegacy",
			args: args{MessageSizeBufferLenLegacy},
			want: 18446744073709551615,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxMessageSize(tt.args.messageSizeBufferLen); got != tt.want {
				t.Errorf("getMaxMessageSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
