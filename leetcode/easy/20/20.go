package main

import "fmt"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
//     左括号必须用相同类型的右括号闭合。
//     左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例 1: 输入: "()" 输出: true
// 示例 2: 输入: "()[]{}" 输出: true
// 示例 3: 输入: "(]" 输出: false
// 示例 4: 输入: "([)]" 输出: false
// 示例 5: 输入: "{[]}" 输出: true

func main() {
	// sSlice := "([)]"
	sSlice := "（【）】"
	fmt.Println(f(sSlice))
}

type stack []rune

func (s *stack) push(ele rune) {
	*s = append(*s, ele)
}

func (s *stack) pop() (ele rune, ok bool) {
	if len(*s) <= 0 {
		return 0, false
	}
	popEle := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popEle, true
}

var matchMap = map[rune]rune{
	'）': '（',
	'】': '【',
	'」': '「',
}

func f(sSlice string) (is bool) {
	st := new(stack)

	for _, val := range sSlice {
		switch val {
		case '（', '【', '「':
			st.push(val)
		case '）', '】', '」':
			if matchVal, ok := st.pop(); matchVal != matchMap[val] || !ok {
				return false
			}
		}
		fmt.Println(string(*st))
	}

	return true
}
