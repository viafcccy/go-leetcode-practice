package main

import (
	"fmt"
	"sort"
)

func main() {
	input1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	input2 := [][]int{{2, 6}, {8, 10}, {1, 3}, {15, 18}}

	fmt.Println(f(input1))
	fmt.Println(f(input2))
}

func f(input [][]int) [][]int {
	sort.Slice(input, func(i, j int) bool { return input[i][0] < input[j][0] })
	res := [][]int{input[0]}
	for _, val := range input[1:] {
		if res[len(res)-1][1] < val[0] {
			res = append(res, val)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], val[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
