package gpt_tasks

import (
	"context"
	"fmt"
	"time"
)


func doWork_(ctx context.Context) {
	for {
		select {
			case <-ctx.Done():
				fmt.Println("Interrupted")
				return
			default:
				fmt.Println("working")
				time.Sleep(50 * time.Millisecond)
		}
	}
}
func Task4() {
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second * 5)
	ctx2, cancel2 := context.WithTimeout(ctx1, time.Second * 2)

	defer cancel1()
	defer cancel2()

	go doWork_(ctx1)
	go doWork_(ctx2)

	<-ctx2.Done()
	fmt.Println("ctx2 Done", ctx2.Err())
	<-ctx1.Done()
	fmt.Println("ctx1 Done", ctx1.Err())

}