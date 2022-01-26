package patternmatcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SinglePrefixMatcher_IsIncluded(t *testing.T) {
	matcher := NewSimplePrefixMatcher("folder-1/b/a")

	assert.True(t, matcher.Matches("folder-1/b/a"))
	assert.True(t, matcher.Matches("folder-1/b/ab"))
	assert.True(t, matcher.Matches("folder-1/b/a/345"))

	assert.False(t, matcher.Matches("folder-1/b"))
}
