package must

import (
	"errors"
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_Must_NoError(t *testing.T) {
	Must(errors.New(nil))
}

func Test_Must_Error(t *testing.T) {
	defer func() {
		r := recover()
		require.NotNil(t, r)
	}()

	Must(errors.New("test error"))
}
