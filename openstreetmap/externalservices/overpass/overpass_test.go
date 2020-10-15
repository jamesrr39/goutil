package overpass

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatBoundsToOverpassFormat(t *testing.T) {
	bounds := Bounds{
		LatMin:  4.14,
		LongMin: -73.73,
		LatMax:  4.21,
		LongMax: -73.67,
	}

	assert.Equal(t, "4.14,-73.73,4.21,-73.67", formatBoundsToOverpassFormat(bounds))
}
