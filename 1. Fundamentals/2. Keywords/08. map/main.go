package main

import "fmt"

/*
=============================
MAP KEYWORD IN GO
=============================

--- 1. Definition of `map` ---
	`map` is a built-in data type in Go that associates keys with values, allowing you to look up a value by its key.
	A `map` is unordered and has unique keys.

--- 2. Purpose of `map` ---
	`map` provides a way to store and retrieve values efficiently based on a key.
	It is commonly used for creating dictionaries, symbol tables, or any structure that requires fast lookups.

--- 3. Syntax ---
	To declare and initialize a map, use the `map` keyword followed by the key and value types:
	```go
	var m map[string]int
	```
	You can also initialize a map using the `make` function:
	```go
	m := make(map[string]int)
	```
	To insert or update a key-value pair:
	```go
	m["key"] = value
	```

--- 4. Accessing and Deleting Elements ---
	- Accessing an element in a map is done using the key:
	```go
	value := m["key"]
	```
	- You can delete an element from a map using the `delete` function:
	```go
	delete(m, "key")
	```

--- 5. Checking Existence of a Key ---
	To check if a key exists in a map, you can use the second value returned by the map access:
	```go
	value, ok := m["key"]
	if ok {
		// key exists
	}
	```

--- 6. Iterating Over Maps ---
	You can iterate over all key-value pairs in a map using a `for` loop:
	```go
	for key, value := range m {
		// key and value can be used here
	}
	```

--- 7. Use Cases ---
	- Storing and managing key-value pairs like dictionaries or lookup tables.
	- Implementing counters, such as counting occurrences of words.
	- Storing configuration settings or user data.

=============================
EXAMPLES OF `MAP` IN GO
=============================

--- 8. Creating a Map ---
	Declaring and initializing a map with `make` and adding key-value pairs.

--- 9. Accessing and Modifying Map Elements ---
	Accessing and modifying map elements using keys.

--- 10. Deleting Elements from a Map ---
	Using the `delete` function to remove an element from a map.

*/

func main() {
	// Example 1: Creating a Map
	m := make(map[string]int)
	m["age"] = 30
	m["height"] = 170
	fmt.Println("Map:", m)

	// Example 2: Accessing and Modifying Map Elements
	if age, exists := m["age"]; exists {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Age not found.")
	}

	m["age"] = 35 // Modifying the value of an existing key
	fmt.Println("Updated Age:", m["age"])

	// Example 3: Deleting Elements from a Map
	delete(m, "height")
	fmt.Println("Map after deletion:", m)

	// Example 4: Iterating Over a Map
	fmt.Println("Iterating over the map:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	// Example 5: Checking Existence of a Key
	if _, exists := m["height"]; exists {
		fmt.Println("Height exists.")
	} else {
		fmt.Println("Height not found.")
	}

	// Example 6: Map with Different Value Types
	student := map[string]interface{}{
		"name":  "Alice",
		"age":   20,
		"grade": 3.5,
	}
	fmt.Println("Student:", student)

	// Example 7: Nested Maps
	grades := map[string]map[string]int{
		"Alice": {"Math": 90, "Science": 85},
		"Bob":   {"Math": 88, "Science": 92},
	}
	fmt.Println("Grades:", grades)

	// Example 8: Map with Slices
	people := map[string][]string{
		"friends": {"Alice", "Bob", "Charlie"},
		"family":  {"David", "Eve"},
	}
	fmt.Println("People:", people)

	// Example 9: Map with Structs
	type Person struct {
		Age  int
		City string
	}
	peopleInfo := map[string]Person{
		"Alice": {Age: 25, City: "New York"},
		"Bob":   {Age: 30, City: "San Francisco"},
	}
	fmt.Println("People Info:", peopleInfo)

	// Example 10: Map of Functions
	funcs := map[string]func(int) int{
		"double": func(x int) int { return x * 2 },
		"triple": func(x int) int { return x * 3 },
	}
	fmt.Println("Functions:", funcs)

	// Example 11: Map of Maps
	contacts := map[string]map[string]string{
		"Alice": {"phone": "123-456", "email": ""},
		"Bob":   {"phone": "789-012", "email": ""},
	}
	fmt.Println("Contacts:", contacts)

	// Example 12: Map of Channels
	channels := map[string]chan int{
		"alice": make(chan int),
		"bob":   make(chan int),
	}
	fmt.Println("Channels:", channels)

	// Example 13: Map of Pointers
	type User struct {
		Name string
		Age  int
	}
	users := map[string]*User{
		"Alice": &User{Name: "Alice", Age: 25},
		"Bob":   &User{Name: "Bob", Age: 30},
	}
	fmt.Println("Users:", users)

	// Example 14: Clear a Map
	for key := range m {
		delete(m, key)
	}

	// or

	m = make(map[string]int) // this will clear the map memory
}
