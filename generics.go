package main

import "fmt"

type cacheable interface {
	int64 | string
}

type Node[T cacheable] struct {
	next *Node[T]
	prev *Node[T]

	data T
}

func NewDLL[T cacheable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

type DoublyLinkedList[T cacheable] struct {
	head *Node[T]
	tail *Node[T]
}

func (L *DoublyLinkedList[cacheable]) Prepend(data cacheable) {
	n := &Node[cacheable]{
		prev: L.head,
		data: data,
	}

	if L.head == nil {
		L.head = n
	} else {
		L.head.next = n
		L.head = n
	}

	if L.tail == nil {
		L.tail = n
	}
}

func (l *DoublyLinkedList[cacheable]) Display() {
	list := l.tail
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}
	fmt.Println()
}

func main() {
	intdll := NewDLL[int64]()

	intdll.Prepend(1)
	intdll.Prepend(2)
	intdll.Prepend(3)
	intdll.Prepend(4)

	intdll.Display()

	strdll := NewDLL[string]()

	strdll.Prepend("a")
	strdll.Prepend("b")
	strdll.Prepend("c")
	strdll.Prepend("d")

	strdll.Display()
}

