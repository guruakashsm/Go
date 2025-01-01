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

/*
InsertAtSpecficPosition inserts a new node at a specific position in the linked list
It takes the data to be inserted and the position at which the new node is to be inserted as arguments
It creates a new node with the given data and inserts it at the given position
It iterates through the list to find the node at the given position
If the position is 0, it inserts the new node at the beginning
If the position is greater than the length of the list, it prints "Index out of Range"
If the position is found, it inserts the new node at the given position
*/
func (list *LinkedList) InsertAtSpecficPosition(data, position int) {
	node := &Node{Data: data, Next: list.Head}
	if position == 0 {
		list.Head = node
		return
	}

	current := list.Head
	var currentPostion int
	for current != nil && position-1 > currentPostion {
		current = current.Next
		currentPostion++
	}

	if current == nil {
		fmt.Println("Index out of Range")
		return
	}

	node.Next = current.Next
	current.Next = node

}

/*
UpdateValueByOldValue updates the value of a node with a given old value in the linked list
It takes the old value and the new value as arguments
It iterates through the list to find the node with the given old value
If the node is found, it updates the value of the node with the new value
If the node is not found, it prints "No data Matched"
*/
func (list *LinkedList) UpdateValueByOldValue(oldValue, newValue int) {
	if list.Head == nil {
		fmt.Println("No data in List")
		return
	}

	current := list.Head

	for current != nil && current.Data != oldValue {
		current = current.Next
	}

	if current == nil {
		fmt.Println("No data Matched")
		return
	}

	current.Data = newValue

}

/*
UpdateAllValueByOldValue updates the value of all nodes with a given old value in the linked list
It takes the old value and the new value as arguments
It iterates through the list to find the nodes with the given old value
If the node is found, it updates the value of the node with the new value
*/
func (list *LinkedList) UpdateAllValueByOldValue(oldValue, newValue int) {
	if list.Head == nil {
		fmt.Println("No data in List")
		return
	}

	current := list.Head

	for current != nil {
		if current.Data == oldValue {
			current.Data = newValue
		}
		current = current.Next
	}

}

/*
UpdateByPosition updates the value of a node at a given position in the linked list
It takes the data and the position as arguments
It iterates through the list to find the node at the given position
If the position is 0, it updates the value of the head node
If the position is greater than the length of the list, it prints "Index out of Range"
If the position is found, it updates the value of the node at the given position
*/
func (list *LinkedList) UpdateByPosition(data, position int) {
	if list.Head == nil {
		fmt.Println("No data Not Found")
		return
	}

	if position == 0 {
		list.Head.Data = data

	}

	current := list.Head
	currentPostion := 0
	for current != nil && position-1 > currentPostion {
		current = current.Next
		currentPostion++
	}

	if current == nil {
		fmt.Println("Index out of Range")
		return
	}

	current.Data = data
}

/*
DeleteByValue deletes the first node with a given value in the linked list
It takes the value to be deleted as an argument
It iterates through the list to find the node with the given value
If the node is found, it deletes the node
If the node is not found, it prints "Value not Found"
*/
func (list *LinkedList) DeleteByValue(data int) {
	if list.Head == nil {
		fmt.Println("No data Found")
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		return
	}

	current := list.Head

	for current.Next != nil && current.Next.Data != data {
		current = current.Next
	}

	if current.Next == nil {
		fmt.Println("Value not Found")
		return
	}

	current.Next = current.Next.Next
}

/*
DeleteByIndex deletes the node at a given index in the linked list
It takes the index of the node to be deleted as an argument
It iterates through the list to find the node at the given index
If the index is 0, it deletes the head node
If the index is greater than the length of the list, it prints "Index out of Range"
If the index is found, it deletes the node at the given index
*/
func (list *LinkedList) DeleteByIndex(index int) {
	if list.Head == nil {
		fmt.Println("No Data Found")
		return
	}

	if index == 0 {
		list.Head = list.Head.Next
		return
	}

	currentPostion := 0
	current := list.Head
	for current.Next != nil && currentPostion < index-1 {
		current = current.Next
		currentPostion++
	}
	if current.Next == nil {
		fmt.Println("Index out of Range")
		return
	}

	current.Next = current.Next.Next

}

/*
DeleteAllByValue deletes all nodes with a given value in the linked list
It takes the value to be deleted as an argument
It iterates through the list to find the nodes with the given value
If the node is found, it deletes the node
*/
func (list *LinkedList) DeleteAllByValue(value int) {
	if list.Head == nil {
		fmt.Println("No Data Found")
		return
	}

	if list.Head.Data == value {
		list.Head = list.Head.Next
	}

	current := list.Head

	for current != nil && current.Next != nil {
		if current.Next.Data == value {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}
}

func main() {
	list := LinkedList{}

	list.InsertAtBack(6)
	list.InsertAtBack(6)
	list.InsertAtFront(2)
	list.InsertAtFront(1)
	list.InsertAfterValue(6, 3)
	list.InsertAfterValue(3, 4)
	// list.InsertInSortedList(7)
	// list.InsertInSortedList(5)
	list.InsertAtSpecficPosition(10, 8)
	list.InsertAtSpecficPosition(11, 9)

	list.UpdateValueByOldValue(6, 2)
	list.UpdateValueByOldValue(5, 6)
	list.UpdateAllValueByOldValue(2, 3)
	list.UpdateAllValueByOldValue(3, 4)
	list.UpdateByPosition(13, 5)
	list.UpdateByPosition(12, 6)

	list.DeleteByValue(4)
	list.DeleteByValue(4)
	list.DeleteByIndex(2)
	list.DeleteByIndex(2)
	list.DeleteAllByValue(3)
	list.DeleteAllByValue(3)

	list.Print()

}
