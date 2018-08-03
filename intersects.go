package redblue_segs

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
)

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func intersects(sa, sb, oa, ob []float64) bool {
	var bln = false
	var a, b, d = segseg_abd(sa[:], sb[:], oa[:], ob[:])

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

func segseg_abd(sa, sb, oa, ob []float64) (float64, float64, float64) {
	var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b float64

	x1, y1 = sa[x], sa[y]
	x2, y2 = sb[x], sb[y]

	x3, y3 = oa[x], oa[y]
	x4, y4 = ob[x], ob[y]

	d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
	a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
	b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))

	return a, b, d
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
