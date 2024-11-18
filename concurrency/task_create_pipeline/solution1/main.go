package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Job struct {
	Value int64
	State State
}

type State int

const (
	InitState State = iota
	FirstStage
	SecondStage
	FinishedStage
)

func main() {
	start := time.Now()
	in := build(50_000_000)
	p1 := task1(in)
	p2 := task2(p1)
	res := task3(p2)

	_ = res

	finish := time.Since(start)
	fmt.Println(finish)
}

func build(iter int) <-chan Job {
	out := make(chan Job)
	go func() {
		for i := 0; i < iter; i++ {
			out <- Job{Value: int64(i)}
		}
		close(out)
	}()
	return out
}

func firstLogic(j *Job) {
	j.Value = int64(float64(j.Value) * math.Pi)
	j.State = FirstStage
}

func secondLogic(j *Job) {
	j.Value = int64(float64(j.Value) * math.E)
	j.State = SecondStage
}

func finishLogic(j *Job) {
	j.Value = int64(float64(j.Value) / float64(rand.Intn(10)))
	j.State = FirstStage
}

func task1(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for task := range in {
			firstLogic(&task)
			out <- task
		}
		close(out)
	}()
	return out
}

func task2(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for task := range in {
			secondLogic(&task)
			out <- task
		}
		close(out)
	}()
	return out
}

func task3(in <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for task := range in {
			finishLogic(&task)
			out <- task
		}
		close(out)
	}()
	return out
}
