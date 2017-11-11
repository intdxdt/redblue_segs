package redblue_segs

import (
	"github.com/intdxdt/math"
)

func prepareEvents(red, blue [][][]float64) []*event {
	var nr = len(red)
	var nb = len(blue)
	var n = nr + nb
	var i, ptr int
	var x, y float64
	var seg [][]float64
	var data = make([]*event, 2*n)

	for i = 0; i < nr; i++ {
		seg = red[i]
		x, y = seg[0][0], seg[1][0]

		data[ptr] = &event{math.MinF64(x, y), CreateRED, i}
		ptr += 1

		data[ptr] = &event{math.MaxF64(x, y), RemoveRED, i}
		ptr += 1
	}
	for i = 0; i < nb; i++ {
		seg = blue[i]
		x, y = seg[0][0], seg[1][0]

		data[ptr] = &event{math.MinF64(x, y), CreateBLUE, i}
		ptr += 1

		data[ptr] = &event{math.MaxF64(x, y), RemoveBLUE, i}
		ptr += 1
	}
	events(data).Sort()
	return data
}
