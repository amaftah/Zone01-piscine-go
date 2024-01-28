package main

import (
	"fmt"

	"piscine"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	node := l.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}

func main() {
	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "2")
	piscine.ListPushFront(link, "you")
	piscine.ListPushFront(link, "man")

	fmt.Println(piscine.ListSize(link))
}
