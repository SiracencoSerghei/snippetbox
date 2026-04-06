package examples

import (
	"context"
	"fmt"
	"time"
)

// Worker функція, яка обробляє завдання
func Worker(ctx context.Context, id int, jobs <-chan int, results chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			// симуляція роботи
			time.Sleep(time.Millisecond * 100)
			results <- fmt.Sprintf("worker %d processed job %d", id, job)
		}
	}
}

// RunWorkers запускає n воркерів і повертає результати
func RunWorkers(ctx context.Context, numWorkers int, numJobs int) []string {
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	// запуск воркерів
	for w := 1; w <= numWorkers; w++ {
		go Worker(ctx, w, jobs, results)
	}

	// відправка завдань
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	var output []string
	for i := 0; i < numJobs; i++ {
		select {
		case res := <-results:
			output = append(output, res)
		case <-ctx.Done():
			return output
		}
	}

	return output
}