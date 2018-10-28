package overpass

import (
	"testing"

	"github.com/jamesrr39/tracks-app/server/domain"
	"github.com/stretchr/testify/assert"
)

func Test_formatBoundsToOverpassFormat(t *testing.T) {
	bounds := &domain.ActivityBounds{
		LatMin:  4.14,
		LongMin: -73.73,
		LatMax:  4.21,
		LongMax: -73.67,
	}

	assert.Equal(t, "4.14,-73.73,4.21,-73.67", formatBoundsToOverpassFormat(bounds))
}
