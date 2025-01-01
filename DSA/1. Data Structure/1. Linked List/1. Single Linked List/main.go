package main

import (
	"fmt"
)

// Node represents a single node in the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a linked list
type LinkedList struct {
	Head *Node
}

/*
Print prints the linked list
It starts from the head and prints the data of each node
It continues until it reaches the end of the list (Next is nil)
*/
func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Println("Data is :", current.Data)
		current = current.Next
	}
}

/*
InsertAtBack inserts a new node at the end of the linked list
It takes the data to be inserted as an argument
It creates a new node with the given data and appends it to the end of the list
First Check if the Head is nil, if yes then assign the node to the Head
If the Head is not nil, then iterate to the end of the list and append the new node
*/
func (list *LinkedList) InsertAtBack(data int) {
	node := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = node
		return
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}

/*
InsertAtFront inserts a new node at the front of the linked list
It takes the data to be inserted as an argument
It creates a new node with the given data and inserts it at the front of the list
If the Head is nil, then assign the node to the Head
If the Head is not nil, then assign the current Head to the Next of the new node and assign the new node to the Head
*/
func (list *LinkedList) InsertAtFront(data int) {
	if list.Head == nil {
		list.Head = &Node{Data: data}
		return
	}

	node := &Node{Data: data, Next: list.Head}
	list.Head = node
}

/*
InsertAfterValue inserts a new node after a given value in the linked list
It takes the value after which the new node is to be inserted and the data to be inserted as arguments
It creates a new node with the given data and inserts it after the node with the given value
It iterates through the list to find the node with the given value
If the node is found, it inserts the new node after it
If the node is not found, it prints "No Record Found"
*/
func (list *LinkedList) InsertAfterValue(afterValue, data int) {
	if list.Head == nil {
		return
	}

	current := list.Head

	for current != nil && current.Data != afterValue {
		current = current.Next
	}

	if current == nil {
		fmt.Println("No Record Found")
		return
	}

	node := &Node{Data: data, Next: current.Next}
	current.Next = node

}

/*
InsertBeforeValue inserts a new node before a given value in the linked list
It takes the value before which the new node is to be inserted and the data to be inserted as arguments
It creates a new node with the given data and inserts it before the node with the given value
It iterates through the list to find the node before the node with the given value
If the node is found, it inserts the new node before it
If the node is not found, it prints "No Record Found"
*/
func (list *LinkedList) InsertBeforeValue(beforeValue, data int) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == beforeValue {
		node := &Node{Data: data, Next: list.Head}
		list.Head = node
		return
	}

	current := list.Head

	for current.Next != nil && current.Next.Data != beforeValue {
		current = current.Next
	}

	if current.Next == nil {
		fmt.Println("No Record Found")
		return
	}

	node := &Node{Data: data, Next: current.Next}
	current.Next = node
}

/*
InsertInSortedList inserts a new node in a sorted linked list
It takes the data to be inserted as an argument
It creates a new node with the given data and inserts it in the sorted list
It iterates through the list to find the correct position to insert the new node
If the list is empty or the data is smaller than the head, it inserts the new node at the beginning
If the data is greater than the head, it iterates through the list to find the correct position to insert the new node
*/
func (list *LinkedList) InsertInSortedList(data int) {
	node := &Node{Data: data, Next: nil}

	if list.Head == nil || list.Head.Data >= data {
		node.Next = list.Head
		list.Head = node
		return
	}

	current := list.Head

	for current.Next != nil && current.Next.Data <= data {
		current = current.Next
	}

	node.Next = current.Next
	current.Next = node
}

func main() {
	list := LinkedList{}

	list.InsertAtBack(3)
	list.InsertAtBack(6)
	list.InsertAtFront(2)
	list.InsertAtFront(1)
	list.InsertAfterValue(3, 5)
	list.InsertBeforeValue(5, 4)
	list.InsertAtBack(9)
	list.InsertAtBack(10)
	list.InsertInSortedList(8)
	list.InsertInSortedList(7)

	list.Print()

}
