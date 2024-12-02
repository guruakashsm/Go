package main

import "fmt"

func main() {
	// 1. Breaking out of a for loop
	fmt.Println("1. Breaking out of a for loop:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop when i is", i)
			break // Exits the loop when i is 5
		}
		fmt.Println(i)
	}

	// 2. Breaking out of an infinite loop
	fmt.Println("\n2. Breaking out of an infinite loop:")
	i := 0
	for {
		i++
		if i > 3 {
			fmt.Println("Breaking out of the infinite loop")
			break // Exits the infinite loop
		}
		fmt.Println(i)
	}

	// 3. Breaking out of a switch statement
	fmt.Println("\n3. Breaking out of a switch statement:")
	value := 2
	switch value {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
		break // Exits the switch statement after case 2
	case 3:
		fmt.Println("Case 3")
	}

	// 4. Breaking out of a labeled loop
	fmt.Println("\n4. Breaking out of a labeled loop:")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of the outer loop")
				break outerLoop // Breaks out of the outer loop
			}
			fmt.Println(i, j)
		}
	}

	// 5. Using break in a switch inside a loop
	fmt.Println("\n5. Using break in a switch inside a loop:")
	for i := 1; i <= 5; i++ {
		switch i {
		case 1:
			fmt.Println("First case")
		case 3:
			fmt.Println("Breaking from loop at i =", i)
			break // Break out of the loop when i is 3
		default:
			fmt.Println("Default case")
		}
	}

	// 6. Breaking out of a labeled loop with a condition
	fmt.Println("\n6. Breaking out of a labeled loop with a condition:")
outerLoop2:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d-%d ", i, j)
			if i == 3 && j == 3 {
				break outerLoop2 // Breaks out of the outer loop
			}
		}
	}
	fmt.Println("\nExited the loop.")

	// 7. Breaking out of a select statement inside a loop
	fmt.Println("\n7. Breaking out of a select statement inside a loop:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "message from ch1"
		ch2 <- "message from ch2"
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
			break // Exits the select block and the loop
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
			break // Exits the select block and the loop
		}
	}
}
