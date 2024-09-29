package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, results chan int) {
	for i := range jobs {
		fmt.Println("wroker", id, "started job", i, "...")
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", i)
		results <- i * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for m := 1; m <= numJobs; m++ {
		<-results
	}
}
