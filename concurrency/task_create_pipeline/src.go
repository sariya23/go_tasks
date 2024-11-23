package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Job - какая-то стуркута, над которой надо проделать операции.
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
	length := 50_000_000
	jobs := make([]Job, length)
	for i := 0; i < length; i++ {
		jobs[i].Value = int64(i)
	}
	start := time.Now()
	jobs = LastProcessing(SecondProcessing(FirstProcessing(jobs)))
	finish := time.Since(start)
	fmt.Println(finish)
}

// FirstProcessing - первый представляет
// первый этап обработки.
func FirstProcessing(jobs []Job) []Job {
	var result []Job
	for _, job := range jobs {
		job.Value = int64(float64(job.Value) * math.Pi)
		job.State = FirstStage
		result = append(result, job)
	}
	return result
}

// SecondProcessing - представляет второй
// этап обработки.
func SecondProcessing(jobs []Job) []Job {
	var result []Job
	for _, job := range jobs {
		job.Value = int64(float64(job.Value) * math.E)
		job.State = SecondStage
		result = append(result, job)
	}
	return result
}

// LastProcessing - представляет последний
// этап обработки.
func LastProcessing(jobs []Job) []Job {
	var result []Job
	for _, job := range jobs {
		job.Value = int64(float64(job.Value) / float64(rand.Intn(10)))
		job.State = FinishedStage
		result = append(result, job)
	}
	return result
}
