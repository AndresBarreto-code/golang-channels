package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d.\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d started fib with %d and answered %d.\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)

}

func main() {
	tasks := []int{2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<-results
	}
}
