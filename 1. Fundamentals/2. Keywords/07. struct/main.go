package main

import "fmt"

/*
=============================
STRUCT IN GO
=============================

--- 1. Definition of `struct` ---
	`struct` is a composite data type in Go that groups together variables (fields) under a single name. Each field in a struct can have a different type, making it a versatile way to represent real-world objects or more complex data structures.

--- 2. Purpose of `struct` ---
	`struct` is used to define and manage related data together as a single unit. It is commonly used to represent an entity with multiple attributes or properties, such as a person, a product, or a car, with each field representing a property of that entity.

--- 3. Syntax of `struct` ---
	```go
	type StructName struct {
		Field1 Type1
		Field2 Type2
		// More fields...
	}
	```
	- `StructName`: The name of the struct type.
	- `Field1`, `Field2`, etc.: The fields of the struct, with each field having a specific type.

--- 4. Creating Struct Instances ---
	Once a struct type is defined, you can create instances of the struct by initializing it either using the named fields or anonymous fields.
	- **Named fields**: Specifying field names during initialization.
	```go
	// Creating a struct instance using named fields
	person := Person{Name: "Alice", Age: 30}
	```
	- **Anonymous fields**: Using positional initialization where field names are not required.
	```go
	// Creating a struct instance using anonymous fields
	point := Point{10, 20}
	```

--- 5. Accessing Struct Fields ---
	To access the fields of a struct, use dot notation (`structInstance.FieldName`).
	```go
	fmt.Println(person.Name)  // Accessing a field named "Name"
	```

--- 6. Pointers to Structs ---
	You can use pointers to structs to avoid copying large structs when passing them to functions or methods. Pointers allow you to modify the original struct inside the function.
	```go
	// Creating a pointer to a struct
	ptr := &person
	// Accessing fields through the pointer
	fmt.Println(ptr.Name)
	```

--- 7. Struct Methods ---
	You can associate methods with structs, which is a key concept in Go's object-oriented programming model. Methods are functions with a receiver type that is the struct type.
	```go
	// Defining a method on a struct
	func (p *Person) Greet() {
		fmt.Println("Hello, my name is", p.Name)
	}
	```

--- 8. Use Cases ---
	- Representing real-world entities like a `Book`, `Person`, or `Car` in an organized manner.
	- Grouping related data together to pass it as a single entity, especially when working with functions or APIs.
	- Defining methods on structs to encapsulate behavior, similar to classes in other programming languages.

=============================
EXAMPLES OF `STRUCT` IN GO
=============================

--- 9. Defining a Simple Struct ---
	Defining a simple struct with basic fields.

--- 10. Using Structs with Methods ---
	Associating methods with structs to define behavior.

--- 11. Struct Pointers ---
	Working with pointers to structs for more efficient memory usage.

*/

type Person struct {
	Name string
	Age  int
}

type Point struct {
	X, Y int
}

// Method on the Person struct
func (p *Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func main() {
	// Example 1: Defining and Using Structs
	fmt.Println("Example 1: Defining and Using Structs")
	person := Person{Name: "Alice", Age: 30}
	fmt.Println("Person:", person)
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)

	// Example 2: Using Structs with Methods
	fmt.Println("\nExample 2: Using Structs with Methods")
	person.Greet()

	// Example 3: Struct Pointers
	fmt.Println("\nExample 3: Struct Pointers")
	ptr := &person
	ptr.Age = 31 // Modify the struct via pointer
	fmt.Println("Updated Person:", *ptr)

	// Example 4: Creating a struct instance using anonymous fields
	point := Point{10, 20}
	fmt.Println("\nPoint:", point)

	// Example 5: Anonymous struct
	var data = struct {
		Name string
		Age  int
	}{Name: "GURU", Age: 25}
	fmt.Println("\nAnonymous Struct:", data)
}
