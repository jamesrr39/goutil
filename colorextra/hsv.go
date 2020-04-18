package colorextra

import (
	"image/color"
	"math"
)

type HSVColor struct {
	H float64 // 0 <= H <= 360
	S float64 // 0 <= S <= 1
	V float64 // 0 <= V <= 1
	A uint8
}

func NewHSVFromRGB(rgb color.RGBA) HSVColor {

	rDash := float64(rgb.R) / float64(255)
	gDash := float64(rgb.G) / float64(255)
	bDash := float64(rgb.B) / float64(255)

	cMax := math.Max(rDash, math.Max(gDash, bDash))
	cMin := math.Min(rDash, math.Min(gDash, bDash))

	// v
	v := cMax

	// s
	var s float64
	if cMax > 0 {
		s = 1 - (cMin / cMax)
	} else {
		s = 0
	}

	return HSVColor{getHue(rgb), s, v, rgb.A}
}
