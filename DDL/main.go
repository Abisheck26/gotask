package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type list struct {
	head *Node
	tail *Node
}

func (l *list) insert(nd int) {
	temp := &Node{nd, nil, nil}
	if l.head == nil {
		l.head = temp
		l.tail = temp
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		temp.prev = p
		p.next = temp
		l.tail = temp
	}
	fmt.Println("Element insereted ", nd)
}

func (l *list) display() {
	p := l.head
	for p != nil {
		fmt.Printf("%d -> ", p.data)
		p = p.next
	}
	fmt.Print("Null")
}

func main() {
	a := list{}
	a.insert(5)
	a.insert(7)
	a.insert(10)
	a.insert(20)
	a.display()

}
