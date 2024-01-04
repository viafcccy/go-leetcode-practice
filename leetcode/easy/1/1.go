package main

import "fmt"

func main() {
	var (
		nums   = []int{7, 1, 2, 15}
		target = 9
	)

	fmt.Println(f(nums, target))
}

func f(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, v1 := range nums {
		if j, ok := m[target-v1]; ok {
			return []int{j, i}
		}

		// 构建一个key是arr val，val是arr index的map。相比暴力循环2遍利用hash表的寻址来节约一次遍历数组。
		m[v1] = i
	}

	return nil
}

// map底层使用hash表存放，寻值时时间复杂度O(1)，类似数组直接索引。
// map类似特殊索引的数组，但是map底层时hash寻址，数组底层是整块的内存空间偏移量寻址
