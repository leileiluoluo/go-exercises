package main

import (
	"errors"
	"fmt"
	"time"
)

type Task struct {
	Id  int
	Err error
	f   func() error
}

func (task *Task) Do() error {
	return task.f()
}

type WorkerPool struct {
	PoolSize    int
	tasksSize   int
	tasksChan   chan Task
	resultsChan chan Task
	Results     func() []Task
}

func NewWorkerPool(tasks []Task, size int) *WorkerPool {
	tasksChan := make(chan Task, len(tasks))
	resultsChan := make(chan Task, len(tasks))
	for _, task := range tasks {
		tasksChan <- task
	}
	close(tasksChan)
	pool := &WorkerPool{PoolSize: size, tasksSize: len(tasks), tasksChan: tasksChan, resultsChan: resultsChan}
	pool.Results = pool.results
	return pool
}

func (pool *WorkerPool) Start() {
	for i := 0; i < pool.PoolSize; i++ {
		go pool.worker()
	}
}

func (pool *WorkerPool) worker() {
	for task := range pool.tasksChan {
		task.Err = task.Do()
		pool.resultsChan <- task
	}
}

func (pool *WorkerPool) results() []Task {
	tasks := make([]Task, pool.tasksSize)
	for i := 0; i < pool.tasksSize; i++ {
		tasks[i] = <-pool.resultsChan
	}
	return tasks
}

func main() {
	t := time.Now()

	tasks := []Task{
		{Id: 0, f: func() error { time.Sleep(2 * time.Second); fmt.Println(0); return nil }},
		{Id: 1, f: func() error { time.Sleep(time.Second); fmt.Println(1); return errors.New("error") }},
		{Id: 2, f: func() error { fmt.Println(2); return errors.New("error") }},
	}
	pool := NewWorkerPool(tasks, 2)
	pool.Start()

	tasks = pool.Results()
	fmt.Printf("all tasks finished, timeElapsed: %f s\n", time.Now().Sub(t).Seconds())
	for _, task := range tasks {
		fmt.Printf("result of task %d is %v\n", task.Id, task.Err)
	}
}
