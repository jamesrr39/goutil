package errorsx

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Error_Error(t *testing.T) {
	err := Errorf("test error created by: %q", "test user")
	assert.Equal(t, "test error created by: \"test user\"", err.Error())
	assert.NotEmpty(t, err.Stack)

	err2 := Wrap(err)
	assert.Equal(t, "test error created by: \"test user\"", err2.Error())
	assert.NotEmpty(t, err2.Stack)

	// stacks should be the same (should take the stack from err)
	assert.Equal(t, err.Stack(), err2.Stack())
}

func Test_Error_Cause_new(t *testing.T) {
	err := errors.New("test error")
	err2 := Wrap(err)
	err3 := Wrap(err2)

	assert.Equal(t, err, Cause(err2))
	assert.Equal(t, err, Cause(err3))
}

func Test_Error_kv(t *testing.T) {
	err := errors.New("test error")
	err = Wrap(err, "k1", "v1", "k2", "v2")
	err = Wrap(err, "k3", "v3")

	assert.Equal(t, `test error [k1="v1", k2="v2", k3="v3"]`, err.Error())
}

func Test_Wrap(t *testing.T) {
	err := Errorf("test error")
	err2 := Wrap(err)
	err3 := Wrap(err2)

	assert.Equal(t, err.Stack(), err2.Stack())
	assert.Equal(t, err2.Stack(), err3.Stack())
}

func TestWrap(t *testing.T) {
	type args struct {
		err     error
		kvPairs []interface{}
	}
	tests := []struct {
		name string
		args args
		want Error
	}{
		{
			"no crash on insufficient args",
			args{
				Errorf("test error"),
				[]interface{}{
					"k1", "v1",
					"k2",
				},
			},
			Errorf(`test error [k1="v1", k2="[empty]"]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.kvPairs...)
			assert.Equal(t, tt.want.Error(), err.Error())
		})
	}
}
