package main

import (
	"fmt"
	"time"
)

func main() {
	// =======================
	// 1. Basic switch statement
	// =======================
	day := 3
	fmt.Println("1. Basic switch statement:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")
	}

	// ==========================
	// 2. Switch with fallthrough
	// ==========================
	fmt.Println("\n2. Switch with fallthrough:")
	month := 5
	switch month {
	case 3:
		fmt.Println("March")
		fallthrough
	case 4:
		fmt.Println("April")
		fallthrough
	case 5:
		fmt.Println("May")
		// Fallthrough keeps going to the next case
	default:
		fmt.Println("Invalid month")
	}

	// ===========================
	// 3. Switch with multiple values
	// ===========================
	fmt.Println("\n3. Switch with multiple values:")
	grade := 85
	switch grade {
	case 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100:
		fmt.Println("Excellent!")
	case 75, 80, 85:
		fmt.Println("Good job!")
	case 60, 65, 70:
		fmt.Println("You passed.")
	default:
		fmt.Println("You failed.")
	}

	// ===========================
	// 4. Select with channels
	// ===========================
	fmt.Println("\n4. Select with channels:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	default:
		fmt.Println("No message received")
	}

	// =======================
	// 5. Default case in select
	// =======================
	fmt.Println("\n5. Default case in select:")
	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	default:
		fmt.Println("No message, default case executed")
	}

	// =======================
	// 6. Unreachable case in select
	// =======================
	fmt.Println("\n6. Unreachable case in select:")
	// This code would give a compilation error due to unreachable case
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2: // Unreachable because it's the same channel
		fmt.Println("Received from ch2:", msg2)
	}
}
