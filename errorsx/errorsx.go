package errorsx

import (
	"fmt"
	"runtime/debug"
	"sort"
	"strings"
)

type kvPairsMapType map[interface{}]interface{}

type Error interface {
	Error() string
	Stack() []byte
}

type Err struct {
	err     error
	kvPairs kvPairsMapType
	stack   []byte
}

func (err *Err) Stack() []byte {
	return err.stack
}

func (err *Err) Error() string {
	var s = err.err.Error()
	var kvStrings []string
	for key, val := range err.kvPairs {
		kvStrings = append(kvStrings, fmt.Sprintf("%s=%q", key, val))
	}
	if len(kvStrings) > 0 {
		sort.Slice(kvStrings, func(i, j int) bool {
			return kvStrings[i] < kvStrings[j]
		})
		s += fmt.Sprintf(" [%s]", strings.Join(kvStrings, ", "))
	}
	return s
}

func Errorf(message string, args ...interface{}) Error {
	return &Err{
		fmt.Errorf(message, args...),
		make(kvPairsMapType),
		debug.Stack(),
	}
}

func Wrap(err error, kvPairs ...interface{}) Error {
	if err == nil {
		return nil
	}

	kvPairsMap := make(kvPairsMapType)
	for i := 0; i < len(kvPairs); i = i + 2 {
		k := kvPairs[i]
		v := kvPairs[i+1]
		kvPairsMap[k] = v
	}

	errType, ok := err.(*Err)
	if !ok {
		return &Err{
			err,
			kvPairsMap,
			debug.Stack(),
		}
	}

	// merge in kv map
	for k, v := range kvPairsMap {
		errType.kvPairs[k] = v
	}

	return errType
}

// Cause fetches the underlying cause of the error
// this should be used with errors wrapped from errors.New()
func Cause(err error) error {
	errErr, ok := err.(*Err)
	if ok {
		return Cause(errErr.err)
	}

	return err
}
