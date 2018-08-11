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

	var redList = createBruteForceList(nr)
	var blueList = createBruteForceList(nb)

	for i := 0; !ret && i < ne; i++ {
		var ev, index = events[i].ev, events[i].idx
		if ev == CreateRED {
			ret = addSegment(index, red, &redList, blue, &blueList, visit, false)
		} else if ev == CreateBLUE {
			ret = addSegment(index, blue, &blueList, red, &redList, visit, true)
		} else if ev == RemoveRED {
			redList.remove(index)
		} else if ev == RemoveBLUE {
			blueList.remove(index)
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
