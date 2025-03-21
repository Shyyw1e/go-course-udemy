package gpt_tasks

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker1(ctx context.Context, wg *sync.WaitGroup, id int) {
	reqID, ok := ctx.Value("requestID").(string)
	userRole, ok1 := ctx.Value("userRole").(string)

	if !ok {
		fmt.Println("Missing requestID in context")
	}
	if !ok1 {
		fmt.Println("Missing userRole in context")
	}
	if !ok || !ok1 {
		wg.Done()
		return
	}

	fmt.Printf("Request ID:\t%s\nUser role:\t%s\n", reqID, userRole)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d interrupted: %s\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d acting . . .\n", id)
			time.Sleep(time.Millisecond * 400)
		}
		
	}
}


func Task7() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	ctx = context.WithValue(ctx, "requestID", "req-123")
	ctx = context.WithValue(ctx, "userRole", "admin")
	defer cancel()

	wg.Add(3)
	go worker1(ctx, wg, 1)
	go worker1(ctx, wg, 2)
	go worker1(ctx, wg, 3)
	
	<-ctx.Done()

	fmt.Println("Waiting . . .")
	wg.Wait()

	fmt.Println("Everything interrupted well")

}