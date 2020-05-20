package snapshottest

import (
	"image"
	"image/color"
	"testing"

	"github.com/jamesrr39/goutil/snapshot"
)

func TestAPISnapshot(t *testing.T) {
	snapshot.AssertMatchesSnapshot(t, "Get /api/example", snapshot.NewTextSnapshot("{\"name\":\"John\"}"))
}

func TestImageSnapshot(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 16, 16))

	img.Set(0, 0, color.Black)

	snapshot.AssertMatchesSnapshot(t, "Image example", snapshot.NewImageSnapshot(img))

}
