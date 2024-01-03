package main

import (
	"fmt"
	"strconv"
)

func main() {
	var target int64 = 0

	fmt.Println(f(target))
	fmt.Println(f2(target))
}

// 整型数字倒置写法 参考002
func f(target int64) bool {
	// 边界条件 结尾是0的数字
	if target < 0 || (target != 0 && target%10 != 0) {
		return false
	}

	var revertedTarget int64
	for target > revertedTarget {
		uintDigit := target % 10
		fmt.Println(uintDigit)
		revertedTarget = revertedTarget*10 + uintDigit
		fmt.Println(revertedTarget)
		target = target / 10
		fmt.Println(target)
	}

	return revertedTarget == target || revertedTarget/10 == target
}

// 字符串写法
func f2(target int64) bool {
	// 熟悉 strconv 库用法
	targetStr := strconv.Itoa(int(target))

	// 熟悉 普通for用法
	for i, j := 0, len(targetStr)/2; i < j; i++ {
		if targetStr[i] != targetStr[len(targetStr)-i] {
			return false
		}
	}

	return true
}
