package main

import (
	"fmt"
	"time"
)

func main() {
	lb := NewLoadBalancer(5)

	lb.AssignTask(func () int {
		time.Sleep(time.Second)
		return 1
	})

	lb.AssignTask(func () int {
		time.Sleep(time.Millisecond * 500)
		return 2
	})

	lb.AssignTask(func () int {
		time.Sleep(time.Millisecond * 750)
		return 3
	})

	lb.AssignTask(func () int {
		time.Sleep(time.Millisecond * 250)
		return 4
	})

	lb.Finished()

	<-lb.Done
}

type Task func () int

func NewLoadBalancer(numOfWorkers int) *LoadBalancer {
	lb := &LoadBalancer{
		tasks:   make(chan Task, 10),
		workers: make([]*TaskWorker, numOfWorkers),
		Done:    make(chan bool, 1),
	}

	for i := 0; i < numOfWorkers; i++ {
		lb.workers[i] = NewTaskWorker(i)
	}

	go lb.AssignmentLoop()

	return lb
}

type LoadBalancer struct {
	roundRobinCounter int
	tasks             chan Task
	workers           []*TaskWorker
	Done              chan bool
}

func (lb *LoadBalancer) AssignTask(task Task) {
	lb.tasks <- task
}

func (lb *LoadBalancer) AssignmentLoop() {
	for {
		task, more := <-lb.tasks

		if task != nil {
			fmt.Printf("Assigning new task to TaskWorker %v.\n", lb.roundRobinCounter)
			lb.workers[lb.roundRobinCounter].Tasks <- task
			lb.roundRobinCounter++
		}

		if !more {
			for _, worker := range lb.workers {
				close(worker.Tasks)
				<-worker.Done
			}

			lb.Done <- true
			return
		}
	}
}

func (lb *LoadBalancer) Finished() {
	close(lb.tasks)
}

func NewTaskWorker(id int) *TaskWorker {
	w := &TaskWorker{
		ID:    id,
		Tasks: make(chan Task, 1),
		Done:  make(chan bool, 1),
	}

	go w.Loop()

	return w
}

type TaskWorker struct {
	ID    int
	Tasks chan Task
	Done  chan bool
}

func (w *TaskWorker) Loop() {
	for {
		task, more := <-w.Tasks

		if task != nil {
			taskID := task()
			fmt.Printf("TaskWorker %v has completed task '%v'.\n", w.ID, taskID)
		}

		if !more {
			w.Done <- true
			return
		}
	}
}
