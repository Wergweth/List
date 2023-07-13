package main

import "fmt"

type Element struct {
	value int
	next  *Element
}
type LinkedList struct {
	head *Element
}

func main() {
	ll := LinkedList{}
	ll.add(111)
	ll.add(222)
	ll.display()
}

func (l *LinkedList) add(value int) {
	NewElement := &Element{value, nil}
	if l.head == nil {
		l.head = NewElement
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = NewElement
	}
}

func (l *LinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("--end--")
}

/*func deleteLastEl(root *Element) {
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
}*/
