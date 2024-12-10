package main

import "fmt"

/*
=============================
FALLTHROUGH IN GO
=============================

--- 1. Definition of `fallthrough` ---
	`fallthrough` is a keyword in Go used in `switch` statements to force the execution to continue 
	to the next case, regardless of whether the case condition is met.

--- 2. Purpose of `fallthrough` ---
	`fallthrough` is used to execute multiple cases in a `switch` statement sequentially. 
	It helps when shared logic or grouped conditions need to be handled without duplicating code.

--- 3. Syntax ---
	Place the `fallthrough` keyword at the end of a `case` block to proceed to the next case:
	```go
	switch value {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
	default:
		fmt.Println("Default case")
	}
	```

--- 4. Behavior of `fallthrough` ---
	- It only works within `switch` statements.
	- It forces execution of the code in the next case, even if its condition is not satisfied.
	- It cannot skip cases; execution proceeds to the immediate next case.

--- 5. Use Cases ---
	- Grouping cases to handle similar logic.
	- Executing a sequence of actions in a `switch` statement.
	- Implementing scenarios where the flow needs to cascade to subsequent cases.

=============================
EXAMPLES OF `FALLTHROUGH` IN GO
=============================

--- 6. Basic `fallthrough` Usage ---
	Using `fallthrough` to execute the next case unconditionally.

--- 7. Grouping Logic ---
	Handling similar cases without repeating code by combining `fallthrough`.

--- 8. Limitations ---
	Understanding that `fallthrough` is only valid for the next immediate case and cannot jump multiple cases.

*/

func main() {
	// Example 1: Basic `fallthrough` usage
	fmt.Println("Basic `fallthrough` example:")
	day := 1
	switch day {
	case 1:
		fmt.Println("Monday")
		fallthrough
	case 2:
		fmt.Println("Tuesday")
		fallthrough
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Other day")
	}

	// Example 2: Grouping logic with `fallthrough`
	fmt.Println("\nGrouping logic example:")
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
		fallthrough
	case "B":
		fmt.Println("Good job!")
		fallthrough
	case "C":
		fmt.Println("You passed.")
	default:
		fmt.Println("Try harder next time.")
	}

	// Example 3: Limitations of `fallthrough`
	fmt.Println("\nLimitations of `fallthrough`:")
	value := 2
	switch value {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
	default:
		fmt.Println("Default case")
	}
}
