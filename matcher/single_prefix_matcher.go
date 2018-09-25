package matcher

import (
	"strings"
)

type SimplePrefixMatcher struct {
	Prefix string
}

func NewSimplePrefixMatcher(prefix string) *SimplePrefixMatcher {
	return &SimplePrefixMatcher{prefix}
}

func (m *SimplePrefixMatcher) IsIncluded(relativePath string) bool {
	return strings.HasPrefix(string(relativePath), m.Prefix)
}
