package concurency

import (
	"fmt"
	"time"
)

func foo(id int, jobs <-chan int, results chan<- int) { // func foo(chName <-chan int) {} только для чтения канала; func foo(chName chan<- int) {} только для записи в канал
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("Worker #%d finished\n", id)
		results<- j * j
	}

}
func WorkerPool() {
	const count = 15
	jobs := make(chan int, count)
	results := make(chan int, count)

	go foo(2, jobs, results)

	for i := range count {
		jobs<-i + 1
	}
	close(jobs)

	for i := range count{
		fmt.Printf("Result # %d:\tvalue = %d\n", i + 1, <-results)
	}
	

}