package gpt_tasks

import (
	"context"
	"fmt"
	"time"
)


func doWork(ctx context.Context) {
	fmt.Print("Working .")
	for i := 0; i < 5; i++ {
		select {
			case <-ctx.Done():
				fmt.Println("\nInterrupted: ", ctx.Err())
				return
			default:
				fmt.Print(". ")
				time.Sleep(time.Second)
		}
		
	}
	fmt.Println("Operation completed")
}

func Task1() {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second * 3)
	defer cancel()

	go doWork(ctx)

	<-ctx.Done()

}