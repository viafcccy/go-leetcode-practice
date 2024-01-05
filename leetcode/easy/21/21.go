package main

import (
	"log"
)

// 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例：输入：1->2->4, 1->3->4 输出：1->1->2->3->4->4

// go build && ./21
func main() {
	// TestLinkedList()

	ll1 := InitLinkedListBySlice([]int{0, 1, 4, 5})
	ll2 := InitLinkedListBySlice([]int{0, 3, 4})

	orderLL := InitLinkedList()

	ll1Pointer := ll1.Head.Next
	ll2Pointer := ll2.Head.Next

	for ll1Pointer != nil && ll2Pointer != nil {
		if ll1Pointer.Value > ll2Pointer.Value {
			orderLL.AddNode(ll2Pointer.Value)
			ll2Pointer = ll2Pointer.Next
		} else {
			orderLL.AddNode(ll1Pointer.Value)
			ll1Pointer = ll1Pointer.Next
		}
	}

	orderLLPointer, err := orderLL.SearchBodyNodeByIndex(uint(orderLL.Len) - 1)
	if err != nil {
		log.Fatalln(err)
	}
	if ll1Pointer != nil {
		orderLLPointer.Next = ll1Pointer.Next
	}
	if ll2Pointer != nil {
		orderLLPointer.Next = ll2Pointer.Next
	}

	orderLL.Print()
}

func TestLinkedList() {
	ll := InitLinkedListBySlice([]int{0, 1, 2, 3, 4})

	node, err := ll.SearchBodyNodeByIndex(3)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(node.Value)

	ll.Print()
}
