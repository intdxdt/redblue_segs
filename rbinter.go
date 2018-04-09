package redblue_segs
const (
	x = iota
	y
)
func RedBlueLineSegmentIntersection(red, blue [][][]float64,
	visit func(int, int) bool) bool {
	var nr = len(red)
	var nb = len(blue)
	var n = nr + nb
	var ne = 2 * n
	var ret bool

	var events = prepareEvents(red, blue)
	//console.log(unpack(events))

	var redList = newBruteForceList(nr)
	var blueList = newBruteForceList(nb)

	for i := 0; i < ne; i++ {
		var ev, index = events[i].ev, events[i].idx
		switch  ev {
		case CreateRED:
			ret = addSegment(index, red, redList, blue, blueList, visit, false)
		case CreateBLUE:
			ret = addSegment(index, blue, blueList, red, redList, visit, true)
		case RemoveRED:
			redList.remove(index)
		case RemoveBLUE:
			blueList.remove(index)
		}
		if ret {
			break
		}
	}

	return ret
}

func RBIntersection(red, blue [][][]float64) [][]int {
	var crossings = make([][]int, 0)
	var visit = func(i, j int) bool {
		crossings = append(crossings, []int{i, j})
		return false
	}

	RedBlueLineSegmentIntersection(red, blue, visit)
	return crossings
}
