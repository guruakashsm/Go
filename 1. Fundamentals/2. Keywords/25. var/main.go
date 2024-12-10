package main

import "fmt"

/*
=============================
VAR IN GO
=============================

--- 1. Definition of `var` ---
	`var` is a keyword in Go used to declare variables. It is used to define a new variable and optionally assign it a value. The variable can be of any type, and Go will infer the type if it is not explicitly stated.

--- 2. Purpose of `var` ---
	`var` is used to declare variables that are meant to store values of different types. This is fundamental to Go’s type system and is used to define variables for general-purpose use, function parameters, return types, and more.

--- 3. Syntax of `var` ---
	```go
	var variableName type
	```
	Where `variableName` is the name of the variable, and `type` is the data type of the variable.

	You can also initialize the variable with a value:
	```go
	var variableName type = value
	```

	Or let Go infer the type based on the value:
	```go
	var variableName = value  // Type is inferred
	```

--- 4. Declaring Multiple Variables ---
	You can declare multiple variables in a single `var` block. The syntax is:
	```go
	var variable1, variable2 type
	```
	You can also initialize multiple variables in a single declaration:
	```go
	var variable1, variable2 = value1, value2
	```

--- 5. Short Variable Declaration (Using `:=`) ---
	In Go, there is a shorthand way of declaring and initializing variables using the `:=` operator. This is useful for local variable declaration inside functions:
	```go
	variableName := value
	```
	Note: `:=` can only be used within functions.

--- 6. Zero Value of Variables ---
	When a variable is declared but not explicitly initialized, Go assigns it a "zero value" based on its type. For example:
	- For numeric types, the zero value is `0`.
	- For strings, the zero value is `""`.
	- For booleans, the zero value is `false`.

--- 7. Use Cases ---
	- Declaring variables to store values.
	- Declaring variables inside functions or globally.
	- Using the zero value for variables that will be initialized later.
	- Utilizing short declarations for local variables inside functions.

=============================
EXAMPLES OF `VAR` IN GO
=============================

--- 8. Declaring a Single Variable ---
	Declaring a single variable with a specific type or initialized value.

--- 9. Declaring Multiple Variables ---
	Declaring multiple variables in a single statement.

--- 10. Zero Value ---
	Declaring variables without initializing them, so they take their zero value.

*/

func main() {
	// Example 1: Declaring a Single Variable
	var age int = 30
	fmt.Println("Age:", age)

	// Example 2: Declaring Multiple Variables
	var name, city string = "Alice", "Wonderland"
	fmt.Println("Name:", name)
	fmt.Println("City:", city)

	// Example 3: Zero Value
	var score int     // Zero value for 'int' is 0
	var isActive bool // Zero value for 'bool' is false
	fmt.Println("Score:", score)
	fmt.Println("Is Active:", isActive)

	// Example 4: Shorthand Variable Declaration
	country := "GoLand" // Using ':=' shorthand for variable declaration
	fmt.Println("Country:", country)
}
