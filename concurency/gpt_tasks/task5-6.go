package gpt_tasks

import (
	"context"
	"fmt"
	"time"
	"sync"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for  {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d interrupted: %s\n", id, ctx.Err())
			
			return
		
		default:
		fmt.Println("Working . . .")
		time.Sleep(time.Millisecond * 300)
		}
	}
}
func Task5() {
	wg := &sync.WaitGroup{}
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()
	
	wg.Add(3)
	go worker(ctx, 1, wg)
	go worker(ctx, 2, wg)
	go worker(ctx, 3, wg)
	
	
	<-ctx.Done()
	fmt.Println("Waiting a bit . . .")
	wg.Wait()
	fmt.Println("All workers finished")
}