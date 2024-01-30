package piscine
package main

import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string

	for n > 0 {
		remainder := n % 10
		result = string('0'+remainder) + result
		n = n / 10
	}

	return result
}


type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head
    for current != nil {
        if comp(current.Data, ref) {
            return &current.Data
        }
        current = current.Next
    }
    return nil

}
