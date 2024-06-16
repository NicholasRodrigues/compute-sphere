package taskmanager

import (
	"errors"
	"sync"
	"time"
)

type TaskManager interface {
	AddTask(task Task) error
	GetTask() (Task, error)
	ProcessTask(task Task) (Result, error)
}

type SumTaskManager struct {
	tasks []Task
	mutex sync.Mutex
}

func (s *SumTaskManager) AddTask(task Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tasks = append(s.tasks, task)
	return nil
}

func (s *SumTaskManager) GetTask() (Task, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.tasks) == 0 {
		return Task{}, errors.New("no tasks available")
	}

	task := s.tasks[0]
	s.tasks = s.tasks[1:]
	return task, nil
}

func (s *SumTaskManager) ProcessTask(task Task) (Result, error) {
	startTime := time.Now() // Start timing

	sum := 0
	for i := 0; i < task.HowMany; i++ {
		sum += task.InitialVal + i
	}

	elapsedTime := time.Since(startTime) // Calculate elapsed time

	result := Result{
		TaskID:         task.ID,
		TimesProcessed: task.TimesToProcess,
		Output:         sum,
		ElapsedTime:    elapsedTime,
	}
	return result, nil
}
