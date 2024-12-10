package main

import "fmt"

/*
=============================
IF KEYWORD IN GO
=============================

--- 1. Definition of `if` ---
	`if` is a keyword in Go that is used for conditional branching. It allows you to execute a block of code
	if a specified condition evaluates to `true`. If the condition evaluates to `false`, the block of code is skipped.

--- 2. Purpose of `if` ---
	`if` is used to make decisions in the program by evaluating a condition. If the condition is `true`,
	the code inside the `if` block is executed. Otherwise, it is skipped, and the program continues.

--- 3. Syntax ---
	The basic syntax of an `if` statement is:
	```go
	if condition {
		// Code to execute if condition is true
	}
	```
	You can also use `else` to handle the scenario when the condition is `false`:
	```go
	if condition {
		// Code for true condition
	} else {
		// Code for false condition
	}
	```

--- 4. `else if` for Multiple Conditions ---
	You can chain multiple conditions using `else if` to check additional conditions if the previous ones fail:
	```go
	if condition1 {
		// Code for condition1
	} else if condition2 {
		// Code for condition2
	} else {
		// Code if none of the conditions are true
	}
	```

--- 5. Use Cases ---
	- Making decisions based on conditions such as user input or variable values.
	- Handling different scenarios based on the outcome of comparisons.
	- Controlling the flow of the program using conditional logic.

--- 6. Best Practices ---
	- Always ensure that the conditions inside `if` statements are clear and simple.
	- Avoid deep nesting of `if` statements; instead, consider using `else if` or `switch` where appropriate.
	- Group related conditions into functions for better readability.

=============================
EXAMPLES OF `IF` IN GO
=============================

--- 7. Basic `if` Usage ---
	Using `if` to execute a block of code when a condition is true.

--- 8. `else if` for Multiple Conditions ---
	Using `else if` to handle multiple conditions.

--- 9. Nested `if` ---
	Using nested `if` statements for more complex decision-making.

*/

func main() {
	// Example 1: Basic `if` usage
	fmt.Println("Basic `if` example:")
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// Example 2: `else if` for multiple conditions
	fmt.Println("\n`else if` example:")
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Example 3: Nested `if`
	fmt.Println("\nNested `if` example:")
	number := 10
	if number >= 0 {
		if number == 0 {
			fmt.Println("Number is zero.")
		} else {
			fmt.Println("Number is positive.")
		}
	} else {
		fmt.Println("Number is negative.")
	}

	// Example 4: Using `if` with logical operators
	fmt.Println("\n`if` with logical operators example:")
	temperature := 30
	isSunny := true
	if temperature > 25 && isSunny {
		fmt.Println("It's a hot and sunny day!")
	} else if temperature <= 25 && isSunny {
		fmt.Println("It's a nice and sunny day.")
	} else {
		fmt.Println("It's a cloudy day.")
	}
}
