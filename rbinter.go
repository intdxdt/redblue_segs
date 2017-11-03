package redblue_segs

func RedBlueLineSegmentIntersection(red, blue [][][]float64,
	visit func(int, int) bool) bool {
	var nr = len(red)
	var nb = len(blue)
	var n = nr + nb
	var ne = 2 * n
	var ret bool

	var events = prepareEvts(red, blue)
	//console.log(unpack(events))

	var redList = NewBruteForceList(nr)
	var blueList = NewBruteForceList(nb)

	for i := 0; i < ne; i++ {
		var ev, index = int(events[i][1]), int(events[i][2])
		switch  ev {
			case CREATE_RED:
				ret = addSegment(index, red, redList, blue, blueList, visit, false)
			case CREATE_BLUE:
				ret = addSegment(index, blue, blueList, red, redList, visit, true)
			case REMOVE_RED:
				redList.remove(index)
			case REMOVE_BLUE:
				blueList.remove(index)
		}
		if ret {
			break
		}
	}

	return ret
}

func RBIntersection(red, blue [][][]float64) [][]int{
    var crossings = make([][]int, 0)
    var visit = func (i, j int) bool{
        crossings = append(crossings, []int{i, j})
        return false
    }

    RedBlueLineSegmentIntersection(red, blue, visit)
    return crossings
}
