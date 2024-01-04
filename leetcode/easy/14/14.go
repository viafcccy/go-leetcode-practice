package main

import (
	"fmt"
	"sort"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1:输入: ["flower","flow","flight"] 输出: "fl"
// 示例 2:输入: ["dog","racecar","car"] 输出: ""
// 解释: 输入不存在公共前缀。
// 说明: 所有输入只包含小写字母 a-z 。

func main() {
	sSlice := []string{"flower", "flow", "flight"}
	fmt.Println(f(sSlice))
}

// 排序后只比较首尾相同的字符
func f(sSlice []string) (res string) {
	switch len(sSlice) {
	case 0:
		return ""
	case 1:
		return sSlice[0]
	}

	// 排序 "" > "a" > ... 着重了解sort包
	sort.Strings(sSlice)

	first := sSlice[0]
	last := sSlice[len(sSlice)-1]
	length := len(first)
	if len(last) < length {
		length = len(last)
	}

	i := 0
	for i < length {
		fmt.Println(i)
		if first[i] != last[i] {
			return first[:i]
		}
		i++
	}

	return
}
