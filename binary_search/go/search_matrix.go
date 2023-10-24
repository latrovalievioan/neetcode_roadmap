//https://leetcode.com/problems/search-a-2d-matrix/
package main

import "fmt"

func main() {
	xs := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3

	fmt.Println(searchMatrix(xs, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	leftRowP, rightRowP := 0, len(matrix)-1

	for leftRowP <= rightRowP {
		mid := (leftRowP + rightRowP) / 2
		if matrix[mid][0] > target {
			rightRowP = mid - 1
		} else if matrix[mid][len(matrix[mid])-1] < target {
			leftRowP = mid + 1
		} else {
			break
		}
	}

	row := matrix[(leftRowP+rightRowP)/2]

	leftColP, rightColP := 0, len(row)-1

	for leftColP <= rightColP {
		mid := (leftColP + rightColP) / 2

		if row[mid] > target {
			rightColP = mid - 1
		} else if row[mid] < target {
			leftColP = mid + 1
		} else {
			return true
		}
	}

	return false
}
