package main

import "fmt"
// List represents a singly-linked list that holds
// values of any type

type List[T any] struct {
	next *List[T]
	val T
}

func main() {
	sn2 := List[string] {nil, "node2"}
	sn1 := List[string] {&sn2, "node1"}
	fmt.Println(sn1, sn2)
	
	nn2 := List[int] {nil, 12}
	nn1 := List[int] {&nn2, 11}
	fmt.Println(nn1, nn2)
}
