package main

import "fmt"

func main() {
	input := "tencent"

	fmt.Println(f(input))
}

func f(input string) int {
	charHash := [26]int{}
	for _, char := range input {
		charHash[char-'a']++
	}

	for i, char := range input {
		if charHash[char-'a'] == 1 {
			return i
		}
	}
	return -1
}
