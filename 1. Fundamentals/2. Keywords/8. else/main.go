package main

import "fmt"

/*
=============================
ELSE IN GO
=============================

--- 1. Definition of `else` ---
	`else` is a keyword in Go that provides an alternative block of code to execute if the `if` condition evaluates to `false`.

--- 2. Purpose of `else` ---
	`else` is used to handle scenarios where the `if` condition is not satisfied, allowing for conditional branching in your code.

--- 3. Syntax ---
	The `else` block follows the `if` block directly without a semicolon or newline:
	```go
	if condition {
		// Code for true condition
	} else {
		// Code for false condition
	}
	```

--- 4. Combining `else` with `if` ---
	You can chain multiple conditions using `else if` to check additional conditions in sequence:
	```go
	if condition1 {
		// Code for condition1
	} else if condition2 {
		// Code for condition2
	} else {
		// Code if none of the conditions are met
	}
	```

--- 5. Use Cases ---
	- Handling alternative logic when the primary condition fails.
	- Implementing decision trees with multiple conditions using `else if`.
	- Providing default logic for cases not handled by `if`.

=============================
EXAMPLES OF `ELSE` IN GO
=============================

--- 6. Basic `else` Usage ---
	Using `else` to handle the alternative scenario when `if` condition is false.

--- 7. `else if` for Multiple Conditions ---
	Chaining conditions using `else if`.

--- 8. Nested `if-else` ---
	Nesting `if-else` blocks to implement more complex logic.

*/

func main() {
	// Example 1: Basic `else` usage
	fmt.Println("Basic `else` example:")
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	// Example 2: `else if` for multiple conditions
	fmt.Println("\n`else if` example:")
	score := 75
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Example 3: Nested `if-else`
	fmt.Println("\nNested `if-else` example:")
	number := -10
	if number >= 0 {
		if number == 0 {
			fmt.Println("Number is zero.")
		} else {
			fmt.Println("Number is positive.")
		}
	} else {
		fmt.Println("Number is negative.")
	}

	// Example 4: Using `else` to provide default logic
	fmt.Println("\nDefault logic with `else`:")
	language := "Python"
	if language == "Go" {
		fmt.Println("You're learning Go!")
	} else if language == "Java" {
		fmt.Println("You're learning Java!")
	} else {
		fmt.Println("Unknown programming language.")
	}
}
