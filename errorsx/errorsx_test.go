package errorsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Error(t *testing.T) {
	err := New("test error")
	assert.Contains(t, string(err.Stack), "\n")
	assert.Equal(t, "", err.Error())
}
