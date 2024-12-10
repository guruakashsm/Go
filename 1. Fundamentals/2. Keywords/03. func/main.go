package main

import "fmt"

/*
=============================
FUNC KEYWORD IN GO
=============================

--- 1. Definition of `func` ---
	`func` is a keyword in Go used to define functions. Functions in Go are blocks of code
	that can be called with specific arguments, perform operations, and return results.

--- 2. Purpose of `func` ---
	`func` is used to define reusable code that can be invoked with or without parameters,
	and optionally return a result. It helps break down complex tasks into smaller, manageable pieces.

--- 3. Syntax ---
	Function declaration follows this basic structure:
	```go
	func functionName(parameters) returnType {
		// Function body
	}
	```
	Example:
	```go
	func add(a int, b int) int {
		return a + b
	}
	```

--- 4. Parameters and Return Values ---
	Functions in Go can accept zero or more parameters and return zero or more values.
	Parameters are specified within parentheses, and the return type follows the parameters.

--- 5. Variadic Functions ---
	Go allows functions to accept a variable number of arguments using the `...` syntax:
	```go
	func sum(numbers ...int) int {
		// Function body
	}
	```

--- 6. Anonymous Functions ---
	Go supports anonymous functions (functions without a name) that can be defined inline and passed around.

--- 7. Use Cases ---
	- Reusable blocks of code.
	- Grouping related operations into logical units.
	- Handling different kinds of tasks like computation, input/output operations, etc.

=============================
EXAMPLES OF `FUNC` IN GO
=============================

--- 8. Basic Function Declaration ---
	Declaring and calling a simple function.

--- 9. Function with Parameters and Return Values ---
	Functions that accept parameters and return results.

--- 10. Variadic Functions ---
	Functions that accept a variable number of arguments.

--- 11. Anonymous Functions ---
	Functions without a name used for specific operations.

*/

func main() {
	// Example 1: Basic function declaration
	fmt.Println("Basic function example:")
	greet()

	// Example 2: Function with parameters and return values
	fmt.Println("\nFunction with parameters and return value example:")
	result := add(10, 20)
	fmt.Println("Result of addition:", result)

	// Example 3: Variadic function example
	fmt.Println("\nVariadic function example:")
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", total)

	// Example 4: Anonymous function example
	fmt.Println("\nAnonymous function example:")
	anonymousFunc := func(message string) {
		fmt.Println(message)
	}
	anonymousFunc("Hello from an anonymous function!")

	// Example 5: Function Direcly Invoked
	func(message string) {
		fmt.Println(message)
	}("Hello from a directly invoked function!")

	// Example 6: Function as a Variable
	var myFunc func(string)
	myFunc = func(message string) {
		fmt.Println(message)
	}
	myFunc("Hello from a function variable!")

	// Example 7: Function as a Parameter
	executeFunction := func(f func(string), message string) {
		f(message)
	}
	executeFunction(func(message string) { fmt.Println(message) }, "Hello from a function parameter!")

}

// Example 1: Basic function declaration
func greet() {
	fmt.Println("Hello, Go!")
}

// Example 2: Function with parameters and return values
func add(a int, b int) int {
	return a + b
}

// Example 3: Variadic function that accepts a variable number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
