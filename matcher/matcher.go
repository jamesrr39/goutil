package matcher

type Matcher interface {
	IsIncluded(relativePath string) bool
}
