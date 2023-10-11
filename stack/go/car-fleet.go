// https://leetcode.com/problems/car-fleet/
// TODO SOLVE AGAIN, why tf is it so slow
package main

import (
	"fmt"
	"sort"
)

func main() {
	// testP := []int{10, 8, 0, 5, 3}
	// testV := []int{2, 4, 1, 1, 3}
	// testTarget := 12
	testP := []int{6, 8}
	testV := []int{3, 2}
	testTarget := 10

	fmt.Println(carFleet(testTarget, testP, testV))
}

func carFleet(target int, position []int, speed []int) int {
	zipped := zip(position, speed)

	sort.SliceStable(zipped, func(i int, j int) bool {
		return zipped[i].p > zipped[j].p
	})

	result := 0
	lastTime := float32(0)

	for i := range zipped {
		arriveTime := ((float32(target) - float32(zipped[i].p)) / float32(zipped[i].v))

		if arriveTime > lastTime {
			lastTime = arriveTime
			result++
		}
	}

	return result
}

type IntZip struct {
	p int
	v int
}

func zip(xs []int, ys []int) []IntZip {
	result := make([]IntZip, len(xs))

	for i := range xs {
		result[i] = IntZip{xs[i], ys[i]}
	}

	return result
}
