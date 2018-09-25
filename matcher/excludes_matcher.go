package matcher

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// ExcludesMatcher is a type that matches file names against excluded names
type ExcludesMatcher struct {
	globs []*regexp.Regexp
}

// NewExcludesMatcherFromReader creates a new ExcludesMatcher from a reader
func NewExcludesMatcherFromReader(reader io.Reader) (*ExcludesMatcher, error) {
	var matcherPatterns []*regexp.Regexp

	bufScanner := bufio.NewScanner(reader)
	for bufScanner.Scan() {
		err := bufScanner.Err()
		if nil != err {
			return nil, err
		}
		pattern := strings.TrimSpace(bufScanner.Text())
		if pattern == "" {
			continue
		}

		if strings.HasPrefix(pattern, "#") {
			// line is a comment
			continue
		}

		matcher, err := regexp.Compile(pattern)
		if nil != err {
			println(pattern)
			return nil, err
		}

		matcherPatterns = append(matcherPatterns, matcher)
	}

	return &ExcludesMatcher{
		globs: matcherPatterns,
	}, nil
}

// Matches tests whether a line matches one of the patterns to be excluded
func (e *ExcludesMatcher) Matches(matched string) bool {
	for _, matcherGlob := range e.globs {
		doesMatch := matcherGlob.MatchString(matched)

		if doesMatch {
			return true
		}
	}
	return false
}
