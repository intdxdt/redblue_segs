package redblue_segs

import (
	"fmt"
	"sort"
	"time"
	"testing"
	"github.com/franela/goblin"
	rb "github.com/intdxdt/redblue_segs/test"
)

var brutal = BruteForce
var rblsi = RBIntersection

func TestCases(t *testing.T) {
	var g = goblin.Goblin(t)

	g.Describe("test cases", func() {
		g.It("should test cases", func() {
			g.Timeout(1 * time.Hour)
			for _, testCase := range rb.Cases {
				fmt.Println(testCase.Name)
				var expected = brutal(testCase.Red, testCase.Blue)
				sort.Sort(rb.Crossings(expected))

				var actual = rblsi(testCase.Red, testCase.Blue)
				sort.Sort(rb.Crossings(actual))
				g.Assert(actual).Equal(expected)

			}
		})
	})
}
