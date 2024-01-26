package main

import "fmt"

func main() {
	input := "abcvba"

	fmt.Println(f(input))
}

func f(input string) bool {
	low, high := 0, len(input)-1
	for low < high {
		if input[low] == input[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if input[i] != input[j] {
					flag1 = false
					break
				}
			}
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if input[i] != input[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
