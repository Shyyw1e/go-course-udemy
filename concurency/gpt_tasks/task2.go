package gpt_tasks

import (
	"context"
	"fmt"
	"time"
)

func Task2() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func(ctx context.Context) {
		for {
			select{
			case <-ctx.Done():
				fmt.Println("Goroutine was stopped")
				return
			default:
				fmt.Println("Working")
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	
	time.Sleep(time.Second * 3)
	cancel()

}