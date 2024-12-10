package main

import "fmt"

/*
=============================
GOTO KEYWORD IN GO
=============================

--- 1. Definition of `goto` ---
	`goto` is a keyword in Go that provides an unconditional jump to another part of the program,
	allowing you to transfer control to a labeled statement. It allows for a non-structured jump
	which can break out of loops, conditionals, or even skip over sections of code.

--- 2. Purpose of `goto` ---
	`goto` is used to transfer control in a program to a specified label, which can be useful
	for breaking out of loops early, or for jumping to certain points in code when necessary.

--- 3. Syntax ---
	The `goto` keyword is used followed by the label name:
	```go
	goto labelName
	```
	Labels are defined by a name followed by a colon `:` before a statement:
	```go
	labelName:
		// Code block here
	```

--- 4. Use Cases ---
	- Breaking out of deeply nested loops.
	- Skipping over sections of code when certain conditions are met.
	- Implementing error handling or cleanup logic in certain situations.

--- 5. Best Practices ---
	While `goto` can be useful in certain situations, it is often considered a "bad practice" because
	it can make code harder to follow and maintain. It is generally advised to use structured control flow
	(loops, conditionals, and functions) when possible.

=============================
EXAMPLES OF `GOTO` IN GO
=============================

--- 6. Basic `goto` Usage ---
	Using `goto` to jump to a specific label in the code.

--- 7. Breaking out of loops with `goto` ---
	Using `goto` to exit from a loop prematurely.

*/

func main() {
	// Example 1: Basic `goto` usage
	fmt.Println("Basic `goto` example:")
	i := 0
startLoop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto startLoop // Jump back to startLoop label
	}

	// Example 2: Breaking out of loops with `goto`
	fmt.Println("\nBreaking out of loop with `goto`:")
	for j := 0; j < 10; j++ {
		if j == 5 {
			goto endLoop // Jump to endLoop label when j reaches 5
		}
		fmt.Println(j)
	}

endLoop:
	fmt.Println("Exited loop early with `goto`.")

	// Example 3: Skipping over code with `goto`
	fmt.Println("\nSkipping code with `goto`:")
	for k := 0; k < 3; k++ {
		if k == 1 {
			goto skipCode // Skip code when k is 1
		}
		fmt.Println("Processing:", k)
	}
skipCode:
	fmt.Println("Code after skipping certain logic.")
}
