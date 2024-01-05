package main

import (
	"bytes"
	"fmt"
	"log"
)

type LinkedListNode struct {
	Value int
	Next  *LinkedListNode
}

type LinkedListHeadNode struct {
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListHeadNode
	Len  int // body length, not include head
}

var (
	ErrNoInit               = fmt.Errorf("linked list do not init")
	ErrCanNotGetHeadNode    = fmt.Errorf("can not get head node")
	ErrCanNotModifyHeadNode = fmt.Errorf("can not modify head node")
	ErrIndexOutRange        = fmt.Errorf("index out range")
	ErrNodeExist            = fmt.Errorf("node exist")
)

func InitLinkedList() (ll *LinkedList) {
	ll = &LinkedList{
		Head: new(LinkedListHeadNode),
		Len:  0,
	}
	return
}

func InitLinkedListBySlice(vales []int) *LinkedList {
	ll := InitLinkedList()
	for _, val := range vales {
		_, err := ll.AddNode(val)
		if err != nil {
			log.Fatal(err)
		}
	}
	return ll
}

func (ll *LinkedList) CheckInit() error {
	if ll.Head == nil {
		return ErrNoInit
	}
	return nil
}

func (ll *LinkedList) CheckEmpty() (bool, error) {
	if err := ll.CheckInit(); err != nil {
		return true, err
	}
	if ll.Head.Next == nil {
		return true, nil
	}
	return false, nil
}

// form first body which is 0
func (ll *LinkedList) SearchBodyNodeByIndex(index uint) (*LinkedListNode, error) {
	if err := ll.CheckInit(); err != nil {
		return nil, err
	}
	if index == 0 {
		return ll.Head.Next, nil
	}
	if index > uint(ll.Len)-1 {
		return nil, ErrIndexOutRange
	}

	currentPointer := ll.Head.Next
	for i := 0; i < int(index); i++ {
		currentPointer = currentPointer.Next
	}

	return currentPointer, nil
}

func (ll *LinkedList) AddNodeInSpecificIndex(val int, index uint) (err error) {
	if err := ll.CheckInit(); err != nil {
		return err
	}
	if index > uint(ll.Len) {
		return ErrIndexOutRange
	}
	if index == 0 {
		ll.Head.Next = &LinkedListNode{
			Value: val,
			Next:  ll.Head.Next,
		}
		ll.Len = ll.Len + 1
		return nil
	}

	lastNode, err := ll.SearchBodyNodeByIndex(index - 1)
	if err != nil {
		return err
	}
	lastNode.Next = &LinkedListNode{
		Value: val,
		Next:  lastNode.Next,
	}
	ll.Len = ll.Len + 1

	return nil
}

func (ll *LinkedList) AddNode(val int) (index uint, err error) {
	err = ll.AddNodeInSpecificIndex(val, uint(ll.Len))
	return uint(ll.Len) - 1, err
}

func (ll *LinkedList) Print() {
	currentPointer := ll.Head.Next

	var buffer bytes.Buffer
	i := 0
	for currentPointer != nil {
		if i == 0 {
			buffer.WriteString("output linked list: ")
			buffer.WriteString("head")
		}
		buffer.WriteString("->")
		buffer.WriteString(fmt.Sprint((*currentPointer).Value))
		currentPointer = (*currentPointer).Next
		i++
	}

	log.Println(buffer.String())
}
