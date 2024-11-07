package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ProccesTasks(tasks []string){
	done := make(chan string)
	completedTasks := []string{}

	for _, task := range tasks {
		go func(task string) {
			sleepTime := time.Duration(rand.Intn(1000)) * time.Millisecond
			time.Sleep(sleepTime)
			done <- task
		}(task)
	}
	for range tasks {
		completedTask := <- done
		completedTasks = append(completedTasks, completedTask)
		fmt.Println(completedTask, "completed")
	}
	fmt.Println("completed tasks:",completedTasks)
	fmt.Println("Alll tasks are done!")
}
func main() {
	rand.Seed(time.Now().UnixNano())
	tasks := []string{"task1", "task2", "task3", "task4", "task5"}
	ProccesTasks(tasks)
}