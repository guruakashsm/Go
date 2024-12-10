package main

import (
	"fmt"
)

/*
=============================
TYPE IN GO
=============================

--- 1. Definition of `type` ---
	`type` is a keyword in Go that allows you to define new types or create aliases for existing types. The `type` keyword is used to introduce user-defined types, which can be based on built-in types or even other user-defined types.

--- 2. Purpose of `type` ---
	`type` is used to improve code readability, create more meaningful names for complex types, and define custom types for better organization and abstraction. It is especially useful when working with complex structures, like struct types, interfaces, or function types.

--- 3. Syntax of `type` ---
	```go
	type TypeName TypeDefinition
	```
	- `TypeName`: The name of the new type.
	- `TypeDefinition`: The type that the new type is based on (e.g., `int`, `struct`, `interface`).

--- 4. Defining Custom Types ---
	You can define a custom type based on existing types (like `int`, `string`, `float64`, etc.) or more complex types such as structs and interfaces.
	```go
	type Age int  // A new type 'Age' based on the built-in type 'int'
	type Name string  // A new type 'Name' based on the built-in type 'string'
	```

--- 5. Creating Type Aliases ---
	You can create a type alias that gives an existing type a new name, which does not introduce a new type but helps with code readability.
	```go
	type MyInt = int  // 'MyInt' is an alias for 'int'
	```

--- 6. Structs and Type ---
	You can define custom types using `struct` to represent complex data structures.
	```go
	type Person struct {
		Name string
		Age  int
	}
	```

--- 7. Type Methods ---
	You can define methods for custom types, similar to how you would for structs. This allows you to encapsulate behavior within your custom types.
	```go
	func (a Age) IsAdult() bool {
		return a >= 18
	}
	```

--- 8. Type Assertion ---
	Go allows you to assert the underlying type of an interface type using type assertions. This can be useful when working with interfaces.
	```go
	var i interface{} = 10
	num, ok := i.(int)  // Type assertion
	```

--- 9. Use Cases ---
	- **Creating Aliases**: For complex or less descriptive types.
	- **Defining Custom Types**: To represent specific entities or concepts, like `Person`, `Employee`, `Account`, etc.
	- **Type Safety**: Creating distinct types to prevent confusion between values of different domains that happen to share the same underlying type.
	- **Method Definition**: Defining methods on custom types to encapsulate behavior.

=============================
EXAMPLES OF `TYPE` IN GO
=============================

--- 10. Defining Custom Types ---
	Creating new types based on built-in types.

--- 11. Type Aliases ---
	Creating type aliases for improved readability.

--- 12. Type Methods ---
	Defining methods for custom types.

*/

type Age int     // A custom type based on int
type Name string // A custom type based on string

// Struct type definition
type Person struct {
	Name string
	Age  Age
}

// Method for Age type
func (a Age) IsAdult() bool {
	return a >= 18
}

func main() {
	// Example 1: Defining Custom Types
	var personAge Age = 25
	var personName Name = "Alice"
	fmt.Println("Name:", personName)
	fmt.Println("Age:", personAge)

	// Example 2: Creating Type Aliases
	type MyInt = int
	var x MyInt = 100
	fmt.Println("MyInt Alias:", x)

	// Example 3: Type Methods
	person := Person{Name: "Bob", Age: 20}
	if person.Age.IsAdult() {
		fmt.Println(person.Name, "is an adult.")
	} else {
		fmt.Println(person.Name, "is not an adult.")
	}

	// Example 4: Type Assertion
	var i interface{} = 10
	num, ok := i.(int)
	if ok {
		fmt.Println("Type Assertion Successful. Value:", num)
	} else {
		fmt.Println("Type Assertion Failed.")
	}

}
