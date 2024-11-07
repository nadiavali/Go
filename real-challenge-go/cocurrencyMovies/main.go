/*Simulate Basic Concurrent Tasks
Challenge: Simulate processing multiple tasks (e.g., "fetching data", "analyzing data") concurrently. 
Each task should take a random time between 100ms and 1 second to complete.
Skills Tested: Concurrency with goroutines and channels.
Hints: Focus on launching goroutines for each task and using channels to receive completion signals.*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Task names to simulate
var tasks = []string{"fetching data", "analyzing data", "saving to database", "sending response"}

// Function to simulate a task with random duration
func processTask(name string, done chan<- string) {
	// Random duration between 100ms and 1 second
	duration := time.Duration(rand.Intn(900)+100) * time.Millisecond
	time.Sleep(duration) // Simulate work by sleeping for the duration
	done <- fmt.Sprintf("Task %s completed in %v", name, duration)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Channel to signal task completion
	done := make(chan string)

	// Launch each task in a separate goroutine
	for _, task := range tasks {
		go processTask(task, done)
	}

	// Wait for each task to complete and print the result
	for range tasks {
		fmt.Println(<-done) // Receive and print the message from the done channel
	}

	fmt.Println("All tasks completed")
}
