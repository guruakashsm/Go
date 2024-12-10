package main

import "fmt"

/*
=============================
RANGE IN GO
=============================

--- 1. Definition of `range` ---
	The `range` keyword in Go is used to iterate over elements in a variety of data structures, such as arrays, slices, maps, and channels. It simplifies looping over these data structures by providing access to both the index (or key) and the value of the element.

--- 2. Purpose of `range` ---
	`range` makes it easy to iterate over data structures in Go, allowing you to access the index and the value of the elements in one step. It is typically used in `for` loops to iterate over arrays, slices, maps, and channels.

--- 3. Syntax of `range` ---
	```go
	for index, value := range collection {
		// Code to execute for each element in the collection
	}
	```

	- `index`: The index (or key for maps) of the current element.
	- `value`: The value of the current element.
	- `collection`: The data structure being iterated over (e.g., array, slice, map, etc.).

--- 4. Range with Different Data Types ---
	- **Arrays and Slices**: When used with arrays and slices, `range` returns the index and the value of each element.
	- **Maps**: When used with maps, `range` returns the key and the value of each entry in the map.
	- **Channels**: When used with channels, `range` iterates over values received from the channel.

--- 5. Ignoring Values ---
	You can ignore the index or value in the `range` loop by using the blank identifier (`_`).

	```go
	for _, value := range collection { // Ignore the index
		// Use the value
	}
	```

	```go
	for index := range collection { // Ignore the value
		// Use the index
	}
	```

--- 6. Use Cases ---
	- Iterating over elements in arrays, slices, maps, or channels.
	- Simplifying loop logic by automatically handling indexing and value extraction.
	- Efficiently receiving values from channels without manually checking for the end condition.

=============================
EXAMPLES OF `RANGE` IN GO
=============================

--- 7. Range on Arrays and Slices ---
	Using `range` to iterate over the elements of an array or slice, accessing both the index and the value.

--- 8. Range on Maps ---
	Using `range` to iterate over the key-value pairs in a map.

--- 9. Range on Channels ---
	Using `range` to iterate over values received from a channel.

*/

func main() {
	// Example 1: Range on an Array
	fmt.Println("Range on Array:")
	arr := [3]int{1, 2, 3}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Example 2: Range on a Slice
	fmt.Println("\nRange on Slice:")
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// Example 3: Range on a Map
	fmt.Println("\nRange on Map:")
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Example 4: Range on Channels (Receiving from a channel)
	fmt.Println("\nRange on Channel:")
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // Closing the channel to stop the loop
	}()

	for value := range ch {
		fmt.Printf("Received from channel: %d\n", value)
	}

	// Example 5: Ignoring the Index
	fmt.Println("\nIgnoring the Index:")
	for _, value := range slice {
		fmt.Printf("Value: %s\n", value)
	}

	// Example 6: Ignoring the Value
	fmt.Println("\nIgnoring the Value:")
	for index := range arr {
		fmt.Printf("Index: %d\n", index)
	}
}
