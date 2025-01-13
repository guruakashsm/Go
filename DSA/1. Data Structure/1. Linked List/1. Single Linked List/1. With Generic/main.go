package main

import (
	"fmt"
)

// Ordered defines a constraint for types that support ordering
type Ordered interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Node represents a single node in the linked list
type Node[T Ordered] struct {
	Data T
	Next *Node[T]
}

// LinkedList represents a linked list
type LinkedList[T Ordered] struct {
	Head *Node[T]
}

/*
Print prints the linked list
It starts from the head and prints the data of each node
It continues until it reaches the end of the list (Next is nil)
*/
func (list *LinkedList[T]) Print() {
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
func (list *LinkedList[T]) InsertAtBack(data T) {
	node := &Node[T]{Data: data, Next: nil}

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
func (list *LinkedList[T]) InsertAtFront(data T) {
	if list.Head == nil {
		list.Head = &Node[T]{Data: data}
		return
	}

	node := &Node[T]{Data: data, Next: list.Head}
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
func (list *LinkedList[T]) InsertAfterValue(afterValue, data T) {
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

	node := &Node[T]{Data: data, Next: current.Next}
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
func (list *LinkedList[T]) InsertBeforeValue(beforeValue, data T) {
	if list.Head == nil {
		return
	}

	if list.Head.Data == beforeValue {
		node := &Node[T]{Data: data, Next: list.Head}
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

	node := &Node[T]{Data: data, Next: current.Next}
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
func (list *LinkedList[T]) InsertInSortedList(data T) {
	node := &Node[T]{Data: data, Next: nil}

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
func (list *LinkedList[T]) InsertAtSpecficPosition(data, position T) {
	node := &Node[T]{Data: data, Next: list.Head}
	if position == T(0) {
		list.Head = node
		return
	}

	current := list.Head
	var currentPostion int
	for current != nil && T(position-1) > T(currentPostion) {
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
func (list *LinkedList[T]) UpdateValueByOldValue(oldValue, newValue int) {
	if list.Head == nil {
		fmt.Println("No data in List")
		return
	}

	current := list.Head

	for current != nil && current.Data != T(oldValue) {
		current = current.Next
	}

	if current == nil {
		fmt.Println("No data Matched")
		return
	}

	current.Data = T(newValue)

}

/*
UpdateAllValueByOldValue updates the value of all nodes with a given old value in the linked list
It takes the old value and the new value as arguments
It iterates through the list to find the nodes with the given old value
If the node is found, it updates the value of the node with the new value
*/
func (list *LinkedList[T]) UpdateAllValueByOldValue(oldValue, newValue int) {
	if list.Head == nil {
		fmt.Println("No data in List")
		return
	}

	current := list.Head

	for current != nil {
		if current.Data == T(oldValue) {
			current.Data = T(newValue)
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
func (list *LinkedList[T]) UpdateByPosition(data, position T) {
	if list.Head == nil {
		fmt.Println("No data Not Found")
		return
	}

	if position == 0 {
		list.Head.Data = data

	}

	current := list.Head
	currentPostion := 0
	for current != nil && position-1 > T(currentPostion) {
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
func (list *LinkedList[T]) DeleteByValue(data int) {
	if list.Head == nil {
		fmt.Println("No data Found")
		return
	}

	if list.Head.Data == T(data) {
		list.Head = list.Head.Next
		return
	}

	current := list.Head

	for current.Next != nil && current.Next.Data != T(data) {
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
func (list *LinkedList[T]) DeleteByIndex(index int) {
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
func (list *LinkedList[T]) DeleteAllByValue(value int) {
	if list.Head == nil {
		fmt.Println("No Data Found")
		return
	}

	if list.Head.Data == T(value) {
		list.Head = list.Head.Next
	}

	current := list.Head

	for current != nil && current.Next != nil {
		if current.Next.Data == T(value) {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}
}

func (list *LinkedList[T]) Length() int {
	n := 0
	current := list.Head
	for current != nil {
		current = current.Next
		n++
	}

	fmt.Println("Count :", n)
	return n

}

func (list *LinkedList[T]) FindIndexByValue(data int) int {
	n := 0
	current := list.Head
	for current != nil && current.Data != T(data) {
		current = current.Next
		n++
	}

	if current == nil {
		return -1
	}

	fmt.Println("Position :", n)
	return n
}

func (list *LinkedList[T]) PrintReverseWithDefer() {
	current := list.Head
	for current != nil {
		defer fmt.Println("Data is :", current.Data)
		current = current.Next
	}
}

func (list *LinkedList[T]) PrintReverseWithRecursion() {

	var printReverse func(node *Node[T])
	printReverse = func(node *Node[T]) {
		if node == nil {
			return
		}
		printReverse(node.Next)
		fmt.Println("Data:", node.Data)
	}

	printReverse(list.Head)
}

func (list *LinkedList[T]) FindMax() {
	if list.Head == nil {
		return
	}

	max := list.Head.Data
	current := list.Head
	for current != nil {
		if current.Data > max {
			max = current.Data
		}
		current = current.Next
	}

	fmt.Println("Max :", max)
}

func (list *LinkedList[T]) FindMin() {
	if list.Head == nil {
		return
	}

	min := list.Head.Data
	current := list.Head
	for current != nil {
		if current.Data < min {
			min = current.Data
		}
		current = current.Next
	}

	fmt.Println("Min :", min)
}

func (list *LinkedList[T]) FindMiddle() {
	if list.Head == nil {
		return
	}

	slow := list.Head
	fast := list.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Printf("Mid Value : %v \n", slow.Data)
}

func (list *LinkedList[T]) Reverse() {
	var prev *Node[T]
	current := list.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}

// TODO : SortList
// TODO : Remove Duplicates

func main() {
	list := LinkedList[int]{}

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

	// list.DeleteByValue(4)
	// list.DeleteByValue(4)
	// list.DeleteByIndex(2)
	// list.DeleteByIndex(2)
	// list.DeleteAllByValue(3)
	// list.DeleteAllByValue(3)
	list.FindIndexByValue(13)
	list.Length()

	// list.PrintReverseWithRecursion()
	list.FindMax()
	list.FindMin()
	list.DeleteAllByValue(4)

	list.Print()
	list.FindMiddle()

}
