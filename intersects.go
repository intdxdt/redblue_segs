package redblue_segs

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
)

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func intersects(sa, sb, oa, ob []float64) bool {
	var bln = false
	var a = ((ob[0] - oa[0]) * (sa[1] - oa[1])) - ((ob[1] - oa[1]) * (sa[0] - oa[0]))
	var b = ((sb[0] - sa[0]) * (sa[1] - oa[1])) - ((sb[1] - sa[1]) * (sa[0] - oa[0]))
	var d = ((ob[1] - oa[1]) * (sb[0] - sa[0])) - ((ob[0] - oa[0]) * (sb[1] - sa[1]))

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	if d == 0 {
		if a == 0.0 && b == 0.0 {
			var o = bbox(oa, ob)
			var box = bbox(sa, sb)
			bln = box.Intersects(o)
		}
		return bln
	}
	//intersection along the the seg or extended seg
	ua := snap_to_zero_or_one(a / d)
	ub := snap_to_zero_or_one(b / d)
	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}

//envelope of segment
func bbox(a, b []float64) *mbr.MBR {
	var box = mbr.CreateMBR(a[x], a[y], b[x], b[y])
	return &box
}

//clamp to zero if float is near zero
func snap_to_zero(v float64) float64 {
	if math.FloatEqual(v, 0.0) {
		v = 0.0
	}
	return v
}

//clamp to zero or one
func snap_to_zero_or_one(v float64) float64 {
	if math.FloatEqual(v, 0.0) {
		v = 0.0
	} else if math.FloatEqual(v, 1.0) {
		v = 1.0
	}
	return v
}
