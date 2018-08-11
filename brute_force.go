package redblue_segs

func BruteForce(red, blue [][][]float64) [][]int {
	var nr = len(red)
	var nb = len(blue)
	var crossings [][]int
	for i := 0; i < nr; i++ {
		for j := 0; j < nb; j++ {
			if intersects(red[i][0], red[i][1],  blue[j][0],  blue[j][1]) {
				crossings = append(crossings, []int{i, j})
			}
		}
	}
	return crossings
}
