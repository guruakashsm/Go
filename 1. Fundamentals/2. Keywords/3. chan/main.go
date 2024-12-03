package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=============================
CHANNELS IN GO
=============================

--- 1. Unbuffered Channels ---
	Unbuffered channels do not have internal storage and block until both sender and receiver are ready.
	Used for synchronous communication between goroutines.

--- 2. Sending Data to Unbuffered Channel ---
	Sending data into an unbuffered channel will block the sender until the receiver is ready to receive the data.

--- 3. Receiving Data from Unbuffered Channel ---
	Receiving from an unbuffered channel will block the receiver until data is available in the channel.

--- 4. Use Case of Unbuffered Channels ---
	Unbuffered channels are useful for synchronizing goroutines, ensuring one goroutine does not continue until another completes.

=============================
BUFFERED CHANNELS
=============================

--- 5. Buffered Channels ---
	Buffered channels allow sending data without blocking as long as there is space in the buffer.
	Once the buffer is full, the sender will block until space becomes available.

--- 6. Creating Buffered Channels ---
	A buffered channel has a fixed capacity defined at creation (e.g., `make(chan int, 2)`).
	Data can be sent to the channel until the buffer is full, after which the sender will block.

--- 7. Receiving Data from Buffered Channels ---
	The receiver can get data from a buffered channel at any time as long as the channel is not empty.

--- 8. Use Case of Buffered Channels ---
	Buffered channels allow asynchronous communication, meaning goroutines can send data into the channel without blocking,
	as long as the buffer is not full.

=============================
CLOSING CHANNELS
=============================

--- 9. Closing a Channel ---
	Closing a channel indicates no more data will be sent.
	It is important to close channels when no more data is being transmitted to signal receivers that no more values are coming.

--- 10. Panic on Sending to Closed Channel ---
	Attempting to send data to a closed channel will cause a runtime panic.
	Always make sure to close channels only when you're sure no more data will be sent.

--- 11. Receiving from a Closed Channel ---
	Receiving from a closed channel returns the zero value for the channel's type when it is empty.
	The second value returned by the receive operation can be used to check if the channel is closed.

=============================
SELECT STATEMENT
=============================

--- 12. `select` Statement ---
	`select` statement allows you to wait on multiple channel operations.
	It is useful for handling multiple concurrent channel operations.

--- 13. Random Selection in `select` ---
	A `select` statement will randomly pick a ready channel if multiple channels are ready for communication.

--- 14. Using `default` Case in `select` ---
	`select` can also include a `default` case, which prevents blocking if no channel is ready.

=============================
GENERAL CHANNELS BEST PRACTICES
=============================

--- 15. Blocking Behavior ---
	Sending data through channels is a blocking operation for unbuffered channels unless there is a receiver available.
	For buffered channels, the sender blocks only when the buffer is full.

--- 16. Channel Safety in Concurrency ---
	Channels in Go provide safe concurrent communication between goroutines without explicit synchronization mechanisms like mutexes.

--- 17. Avoid Sending to Closed Channels ---
	It's important to avoid sending data to a closed channel, as it will result in a panic. Ensure proper synchronization when closing channels.

--- 18. Passing Channels to Functions ---
	Channels can be passed as arguments to functions, and you can restrict whether a channel can only send (`chan<-`) or only receive (`<-chan`) data.

--- 19. Handling Deadlocks ---
	Always ensure that data sent to channels is properly handled. If a channel is not read from or written to correctly, it can cause deadlocks or data loss.

Why to use Channels instead of Variables?
	Channels in Go provide a safer, more idiomatic way to communicate between goroutines by automatically synchronizing access to shared data, preventing data races, and avoiding the need for explicit locking mechanisms (e.g., using sync.Mutex).
	They allow goroutines to pass data in a structured, orderly way, ensuring safe concurrent operations without the complexities and risks associated with manually handling shared variables.

*/

// Worker pool using channels
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark the goroutine as done when it finishes
	for job := range jobs {
		fmt.Printf("Worker %d started processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}
}

// Fan-Out example: One goroutine sending data to multiple receivers (goroutines)
func fanOutExample() {
	// Create a shared channel to send data to multiple goroutines
	ch := make(chan int)

	var wg sync.WaitGroup

	// Worker function that processes data from the channel
	worker := func(id int) {
		defer wg.Done()       // Mark this goroutine as done once it finishes processing
		for job := range ch { // Read from the channel
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Second) // Simulate work being done
		}
	}

	// Launch 3 worker goroutines that will receive from the channel
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Send 5 jobs to the channel (One sender)
	for i := 0; i < 5; i++ {
		ch <- i
	}

	// Close the channel after sending all jobs to signal workers that no more data is coming
	close(ch)

	// Wait for all worker goroutines to finish processing
	wg.Wait()
}

// Fan-In example: Multiple goroutines sending data to a single channel
func fanInExample() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// 3 worker goroutines pushing data to the channel
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				ch <- i*10 + j
			}
		}(i)
	}

	// Receiving from the channel
	go func() {
		wg.Wait() // Wait for all workers to finish
		close(ch) // Close the channel when all workers are done
	}()

	for msg := range ch {
		fmt.Println("Fan-In received:", msg)
	}
}

// Using Select to handle multiple channels
func selectExample() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 42
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 99
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}
}

// Channel as a signal (done channel)
func doneSignalExample() {
	done := make(chan bool)
	go func() {
		// Simulate work
		time.Sleep(2 * time.Second)
		done <- true
	}()
	<-done
	fmt.Println("Work is done!")
}

// Buffered channel example
func bufferedChannelExample() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Buffered Channel:", <-ch, <-ch, <-ch)
}

// Using channels with timeout
func timeoutExample() {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()
	select {
	case val := <-ch:
		fmt.Println("Received:", val)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout!")
	}
}

// Closing a channel
func closeChannelExample() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel
	}()
	for val := range ch {
		fmt.Println("Received from closed channel:", val)
	}
}

func main() {
	// Show worker pool usage
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs to the workers
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results
	wg.Wait() // Ensure that all workers have finished before collecting results
	for i := 0; i < 5; i++ {
		fmt.Println("Worker result:", <-results)
	}

	// Example of Fan-Out, Fan-In, Select, and other channel patterns
	fmt.Println("\nFan-Out Example:")
	fanOutExample()

	fmt.Println("\nFan-In Example:")
	fanInExample()

	fmt.Println("\nSelect Example:")
	selectExample()

	fmt.Println("\nDone Signal Example:")
	doneSignalExample()

	fmt.Println("\nBuffered Channel Example:")
	bufferedChannelExample()

	fmt.Println("\nTimeout Example:")
	timeoutExample()

	fmt.Println("\nClose Channel Example:")
	closeChannelExample()
}
