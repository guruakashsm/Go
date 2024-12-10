package main

import "fmt"

/*
=============================
CONTINUE IN GO
=============================

--- 1. Definition of `continue` ---
	The `continue` statement is used to skip the remaining part of the current iteration in a loop and proceed with the next iteration.
	It helps to avoid unnecessary code execution when certain conditions are met.

--- 2. How `continue` Works ---
	When a `continue` statement is encountered, it immediately stops the current iteration of the loop and moves to the next iteration.
	This is useful when you want to skip specific conditions and continue processing the next values.

--- 3. Common Use Case of `continue` ---
	`continue` is commonly used to skip over unwanted iterations, such as when filtering data or avoiding unnecessary calculations.

=============================
EXAMPLES OF `CONTINUE` IN GO
=============================

--- 4. Basic Usage of `continue` ---
	Using `continue` in a loop to skip specific conditions.

--- 5. Skipping Odd Numbers ---
	Using `continue` to skip odd numbers in a loop.

--- 6. Skipping Even Numbers ---
	Using `continue` to skip even numbers in a loop.

--- 7. Nested Loops and `continue` ---
	Using `continue` in nested loops to skip the inner loop iteration.

*/

func main() {
	// Example 1: Skipping odd numbers using continue
	fmt.Println("Skipping odd numbers:")
	for i := 1; i <= 10; i++ {
		// Skip odd numbers
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// Example 2: Skipping even numbers using continue
	fmt.Println("\nSkipping even numbers:")
	for i := 1; i <= 10; i++ {
		// Skip even numbers
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Example 3: Nested loops with continue
	fmt.Println("\nNested loops with continue:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			// Skip when i == j
			if i == j {
				continue
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	// Example 4: Labeling with continue
	fmt.Println("\nNested loops with Labeling continue:")
firstloop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			// Skip when i == j
			if i == j {
				continue firstloop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
