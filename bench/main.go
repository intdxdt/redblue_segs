package main

import (
	"time"
	"fmt"
	rb "github.com/intdxdt/redblue_segs"
	rbt "github.com/intdxdt/redblue_segs/test"
	"github.com/intdxdt/math"
)

const (
	NUM_ITER = 100
	WARM_UP  = 10
)
var IntersectBruteForce = rb.BruteForce

type Task struct {
	name string
	algo algorFn
}
type algorFn func([][][]float64, [][][]float64) [][]int

var implementations = []Task{
	{name: "Brute-force",   algo: IntersectBruteForce},
	{name: "RBLSI",         algo: rb.RBIntersection},
}

func benchmark(red, blue [][][]float64, algo algorFn) []float64 {
	for i := 0; i < WARM_UP; i++ {
		algo(red, blue)
	}

	var start = time.Now()
	var count = 0

	for i := 0; i < NUM_ITER; i++ {
		result := algo(red, blue)
		count += len(result)
	}
	var end = time.Now()
	return []float64{math.Round((end.Sub(start).Seconds() * 1000.0) / float64(NUM_ITER),2), float64(count)}
}

func main () {
	for i := 0; i < len(implementations); i++ {
		impl := implementations[i]
		fmt.Println("testing", impl.name, "...")
		for j := 0; j < len(rbt.Cases); j++ {
			fmt.Println("case:", rbt.Cases[j].Name)
			res := benchmark(rbt.Cases[j].Red, rbt.Cases[j].Blue, impl.algo)
			fmt.Println(fmt.Sprintf("\ttime: %v ms - total isect: %v", res[0], res[1]))
		}
	}
}
