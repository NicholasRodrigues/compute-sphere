package main

import (
	"fmt"
	"github.com/NicholasRodrigues/compute-sphere/server/internal/taskmanager"
	"sync"
	"time"
)

func worker(id int, tasks <-chan taskmanager.Task, results chan<- taskmanager.Result, tm taskmanager.TaskManager, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result, err := tm.ProcessTask(task)
		if err != nil {
			continue
		}
		results <- result
	}
}

func main() {
	const numWorkers = 1000000
	const numTasks = 1000000

	tasks := make(chan taskmanager.Task, numTasks)
	results := make(chan taskmanager.Result, numTasks)
	finalResult := make(chan int)

	var wg sync.WaitGroup

	tm := &taskmanager.SumTaskManager{}

	for i := 1; i <= numTasks; i++ {
		task := taskmanager.Task{
			ID:             i,
			Name:           fmt.Sprintf("SumTask %d", i),
			InitialVal:     1,
			HowMany:        10,
			TimesToProcess: 1,
		}
		tm.AddTask(task)
	}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, tm, &wg)
	}

	overallStartTime := time.Now()

	go func() {
		totalSum := 0
		for result := range results {
			totalSum += result.Output
			fmt.Printf("Result of task %d: Sum = %d, Elapsed time = %v\n", result.TaskID, result.Output, result.ElapsedTime)
		}
		finalResult <- totalSum
	}()

	go func() {
		for {
			task, err := tm.GetTask()
			if err != nil {
				break
			}
			tasks <- task
		}
		close(tasks)
	}()

	wg.Wait()
	close(results)

	totalSum := <-finalResult

	overallElapsedTime := time.Since(overallStartTime)
	fmt.Printf("Overall elapsed time: %v\n", overallElapsedTime)
	fmt.Printf("Final sum of all results: %d\n", totalSum)
}
