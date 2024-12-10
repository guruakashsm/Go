package main

import "fmt"

/*
=============================
CONSTANTS IN GO
=============================

--- 1. Declaring Constants ---
	Constants in Go are values that cannot be changed after they are assigned.
	They are defined using the `const` keyword followed by the name of the constant and its value.

--- 2. Constant Types ---
	Constants can be of any basic type: numeric, string, or boolean.
	You can declare multiple constants in a single `const` block.

--- 3. Immutability ---
	Once a constant is assigned a value, that value cannot be changed, and attempting to modify it will result in a compilation error.

--- 4. Typed and Untyped Constants ---
	Constants can be typed (like `const x int = 5`) or untyped (like `const x = 5`), where Go infers the type.

--- 5. Use Case of Constants ---
	Constants are often used for fixed values that are reused in the program, such as configuration settings, mathematical constants, or status codes.

--- 6. iota Keyword ---
	The `iota` keyword is often used in constants when you need to create a sequence of related constants, especially useful for enumerations.

=============================
EXAMPLES OF CONSTANTS IN GO
=============================

--- 7. Declaring Constants ---
	Constants are declared using the `const` keyword followed by the name, type, and value.
	For example, `const Pi = 3.14159` declares a constant for the value of Pi.

--- 8. Using `iota` for Sequence ---
	`iota` allows you to generate sequential values, often used for enumerating sets of constants.

=============================
BEST PRACTICES WITH CONSTANTS
=============================

--- 9. Grouping Constants ---
	You can group constants together using a `const` block for better organization.

--- 10. Naming Conventions ---
	Follow proper naming conventions for constants. By convention, constants in Go are usually written in all uppercase letters with underscores separating words (e.g., `MAX_RETRIES`).

*/

const (
	// Declaring constants
	PI        = 3.14159
	STATUS_OK = 200
	STATUS_NOT_FOUND = 404
)

const (
	// Using iota to create a sequence of constants
	Sunday = iota
	Monday
	Tuesday
)

func main() {
	// Accessing constants
	fmt.Println("PI constant:", PI)
	fmt.Println("Status OK:", STATUS_OK)
	fmt.Println("Status NOT FOUND:", STATUS_NOT_FOUND)

	// Using iota for sequence
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)

	// Constants cannot be modified
	// Uncommenting the following line will cause a compilation error
	// PI = 3.14 // Error: cannot assign to PI

	// Demonstrating typed and untyped constants
	const X = 10 // Untyped constant
	const Y int = 20 // Typed constant
	fmt.Println("X:", X)
	fmt.Println("Y:", Y)

	// Multiple constants in a single declaration
	const (
		MinAge = 18
		MaxAge = 100
	)
	fmt.Println("Min Age:", MinAge)
	fmt.Println("Max Age:", MaxAge)
}

