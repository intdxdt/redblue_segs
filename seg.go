package redblue_segs

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/segs"
)

func addSegment(
	index int, red [][][]float64, redList *BruteForceList,
	blue [][][]float64, blueList *BruteForceList,
	visit func(int, int) bool, flip bool) bool {
	//Look up segment
	var seg = red[index]

	//Get segment end points
	var x0 = seg[0]
	var x1 = seg[1]

	//Read out components
	var a0 = x0[1]
	var a1 = x1[1]
	var l0 = math.MinF64(a0, a1)
	var h0 = math.MaxF64(a0, a1)

	//Scan over blue intervals for point
	var intervals = blueList.intervals
	var blueIndex = blueList.index
	var count = blueList.count
	var ptr = 2 * count

	for i := count - 1; i >= 0; i-- {
		ptr += -1
		var h1 = intervals[ptr]
		ptr += -1
		var l1 = intervals[ptr]

		//Test if intervals overlap
		if l0 <= h1 && l1 <= h0 {
			var bindex = blueIndex[i]
			var bseg = blue[bindex]

			//Test if segments intersect
			if segs.Intersects(seg, bseg) {
				var ret bool
				if flip {
					ret = visit(bindex, index)
				} else {
					ret = visit(index, bindex)
				}
				if ret {
					return ret
				}
			}
		}
	}

	redList.insert(l0, h0, index)
	return false
}
