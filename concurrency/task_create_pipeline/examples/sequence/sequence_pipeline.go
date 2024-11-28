package main

import (
	"fmt"
	"time"
)

type Job func(in <-chan int) <-chan int

func InitPipelineData() <-chan int {
	out := make(chan int)

	go func() {
		for i := range 10 {
			out <- i
		}
		close(out)
	}()

	return out
}

func FirstPipelineJob(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			time.Sleep(time.Second)
			// Тут непосредственно действие, которое совершаем над данныме во входном канале.
			// При этом параллелим эти действия в отдельных горутинах.
			out <- i * 2
		}
		close(out)
	}()

	return out
}

func SecondPipelineJob(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			time.Sleep(time.Second)
			// Тут непосредственно действие, которое совершаем над данныме во входном канале.
			// При этом параллелим эти действия в отдельных горутинах.
			out <- i + 1
		}
		close(out)
	}()

	return out
}

func ThirdPipelineJob(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for i := range in {
			time.Sleep(time.Second)
			// Тут непосредственно действие, которое совершаем над данныме во входном канале.
			// При этом параллелим эти действия в отдельных горутинах.
			out <- i * 10
		}
		close(out)
	}()

	return out
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Итоговое время работы: %s", time.Since(start))
	}()

	jobs := []Job{
		Job(FirstPipelineJob),
		Job(SecondPipelineJob),
		Job(ThirdPipelineJob),
	}

	result := InitPipelineData()
	for _, job := range jobs {
		result = job(result)
	}

	for data := range result {
		fmt.Printf("Результат после прохождения всех шагов пайплайна: %d\n", data)
	}
}
