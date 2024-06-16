package taskmanager

import "time"

type Tasks []Task
type Task struct {
	ID             int
	Name           string
	Payload        string
	InitialVal     int
	HowMany        int
	TimesToProcess int
}

type Result struct {
	TaskID         int
	TimesProcessed int
	Output         int
	ElapsedTime    time.Duration
}
