package snapshottest

import (
	"testing"

	"github.com/jamesrr39/goutil/snapshot"
)

func TestAPISnapshot(t *testing.T) {
	snapshot.AssertMatchesSnapshot(t, "Get /api/example", "{\"name\":\"John\"}")
}
