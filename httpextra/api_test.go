package httpextra

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeJSONDataResponse(t *testing.T) {
	type T struct {
		ID        int64
		CreatedAt time.Time `json:"createdAt"`
		UAString  string    `json:"uaString"`
	}

	bb := bytes.NewBufferString(`{"data":{"ID":3,"createdAt":"1970-01-01T00:00:01.000000Z","uaString":"Mozilla"}}`)

	tt := new(T)
	err := DecodeJSONDataResponse(bb, &tt)
	require.NoError(t, err)

	assert.Equal(t, "Mozilla", tt.UAString)
	assert.Equal(t, time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC), tt.CreatedAt)
}
