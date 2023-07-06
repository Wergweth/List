package main

import "fmt"

type Element struct {
	value int
	next  *Element
}

func main() {
	el3 := Element{
		4,
		nil,
	}
	el2 := Element{
		3,
		&el3,
	}
	el1 := Element{
		2,
		&el2,
	}
	el0 := Element{
		1,
		&el1,
	}
	printList(&el0)
	reverse(&el0)
	printList(&el3)
}

func printList(root *Element) {
	if root == nil {
		fmt.Println("root nil")
		return
	}
	fmt.Println(root.value)
	if root.next != nil {
		printList(root.next)
	} else {
		fmt.Println("end")
		return
	}
}

func deleteLastEl(root *Element) {
	if root == nil {
		fmt.Println("root nil")
		return
	}
	if root.next.next == nil {
		root.next = nil
	} else {
		deleteLastEl(root.next)
	}
}

func addLastEl(newVal int, root *Element) {
	if root == nil {
		fmt.Println("root nil")
		return
	}
	if root.next == nil {
		root.next = &Element{newVal, nil}
		return
	} else {
		addLastEl(newVal, root.next)
	}
}

func reverse(root *Element) {
	if root == nil {
		fmt.Println("root nil")
		return
	}
	var current *Element
	var next *Element
	for root != nil {
		current = root.next
		root.next = next
		next = root
		root = current
	}
	return
}
