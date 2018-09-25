package humanise

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HumaniseBytes(t *testing.T) {
	assert.Equal(t, "1000.0 B", HumaniseBytes(1000))
	assert.Equal(t, "1.0 KiB", HumaniseBytes(1024))
	assert.Equal(t, "3.5 MiB", HumaniseBytes(1024*1024*3.5))
	assert.Equal(t, "8.0 GiB", HumaniseBytes(1024*1024*1024*8))

	bigSize := math.Pow(2, 62) * float64(1.5)
	assert.Equal(t, "6.0 EiB", HumaniseBytes(int64(bigSize)))
}
