package main

/*
    Program Description:
    This Go program demonstrates concurrent task processing using goroutines, channels, and sync.WaitGroup.

    Components:
    - **Channels**: A buffered channel `work` is used to pass tasks (strings) to the worker goroutines.
    - **Worker Goroutines**: 5 worker goroutines are launched, each processing tasks concurrently from the channel.
    - **sync.WaitGroup**: Ensures that the main function waits for all worker goroutines to complete their work before exiting.
    - **Task Distribution**: 15 predefined tasks (like "Cleaning", "Washing", etc.) are sent to the `work` channel. 
      Each worker picks up tasks, logs the task, and simulates work by sleeping for 1 second.
    - **Channel Closing**: Once all tasks are sent, the channel is closed to signal workers to stop reading from it.

    Workflow:
    1. The main function initializes the `work` channel and the `sync.WaitGroup`.
    2. 5 worker goroutines are started, and they wait for tasks to be sent through the `work` channel.
    3. The main function sends 15 tasks (from an array) to the channel.
    4. Once all tasks are sent, the channel is closed to notify workers that no more tasks will be available.
    5. The `sync.WaitGroup` ensures that the main function waits for all workers to finish processing before it exits.

    Key Concepts:
    - Goroutines for concurrency.
    - Channels for communication between goroutines.
    - WaitGroup for synchronization and ensuring all tasks are completed before exiting.
    - Buffered channels to allow sending multiple tasks without blocking the main function.

    Note:
    - If more tasks need to be processed, increase the number of tasks in the array and adjust the number of workers accordingly.
    - Proper handling of the channel closing is essential to avoid any goroutine leaks or panics.

    Usage:
    - This structure is useful for scenarios where concurrent task processing is needed, such as processing jobs in parallel or handling multiple I/O-bound tasks simultaneously.
*/


import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	work := make(chan string, 10)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(work, i, &wg)
	}
	arr := []string{
		"Cleaning", "Washing", "Watching",
		"Reading", "Writing", "Running",
		"Cooking", "Baking", "Frying",
		"Swimming", "Diving", "Snorkeling",
		"Traveling", "Hiking", "Camping",
		"Painting", "Sketching", "Drawing",
		"Cycling", "Jogging", "Walking",
		"Shopping", "Browsing", "Purchasing",
		"Listening", "Composing", "Playing",
		"Organizing", "Tidying", "Cleaning",
		"Studying", "Learning", "Researching",
		"Meditating", "Relaxing", "Breathing",
		"Exploring", "Discovering", "Adventuring",
		"Building", "Constructing", "Assembling",
		"Dancing", "Singing", "Acting",
	}

	for i := 0; i < 15; i++ {
		work <- arr[i]
	}
	close(work)
	wg.Wait()
}

func worker(works <-chan string, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for work := range works {
		log.Printf("Worker %v started doing %v", i, work)
		time.Sleep(1 * time.Second)
	}

}
