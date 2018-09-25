package matcher

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewExcludesMatcherFromReader(t *testing.T) {
	buf := bytes.NewBuffer([]byte(`
# comment
.caches/*
^.*.mp4
.android/*

`))

	matcher, err := NewExcludesMatcherFromReader(buf)
	require.Nil(t, err)

	assert.Len(t, matcher.globs, 3, "expected 2 matcher patterns - has the comment or blank been included as a regex")

	assert.True(t, matcher.Matches("a/b/myvideo.mp4"))

	assert.True(t, matcher.Matches(".caches/a.txt"))

	assert.True(t, matcher.Matches(".android/avd/Nexus_5_API_22.avd/system.img.qcow2"))

	assert.False(t, matcher.Matches("a/b/mypic.jpg"))

}
