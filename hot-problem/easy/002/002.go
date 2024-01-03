package main

import "math"

func main() {
	var target int32 = -123

	print(f(target))
}

func f(target int32) (result int) {

	// 注意边界处理
	if target == 0 {
		return 0
	}

	for target != 0 {

		// 取出个位数
		unitDigit := target % 10
		result = result*10 + int(unitDigit)

		// 注意边界处理
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}

		// 舍弃精度
		target = target / 10
	}

	return
}
