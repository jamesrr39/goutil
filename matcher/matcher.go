package matcher

import "github.com/jamesrr39/intelligent-backup-store-app/intelligentstore/intelligentstore"

// FIXME: rename package

type Matcher interface {
	IsIncluded(relativePath intelligentstore.RelativePath) bool
}
