package gpt_tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker2(ctx context.Context, wg *sync.WaitGroup, id int) {
	requestID, ok1 := ctx.Value("requestID").(string)
	userRole, ok2 := ctx.Value("userRole").(string)
	defer wg.Done()
	if !ok1 {
		fmt.Println("No request id")
		return
	}
	if !ok2 {
		fmt.Println("No user role")
		return
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d interrupted\n", id)
			return
		default:
			fmt.Printf("Worker %d working\nRequest ID:\t%s\nUser Role:\t%s", id, requestID, userRole)
			time.Sleep(time.Millisecond*400)
		}
	}
}


func Task8() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	ctx = context.WithValue(ctx, "requestID", "req-123")
	ctx = context.WithValue(ctx, "userRole", "admin")
	defer cancel()

	signCh := make(chan os.Signal, 3)
	signal.Notify(signCh, os.Interrupt)

	wg.Add(3)
	go worker2(ctx, wg, 1)
	go worker2(ctx, wg, 2)
	go worker2(ctx, wg, 3)

	<-signCh
	fmt.Println("SIGINT received. Shutting down gracefully...")
	
	cancel()

	fmt.Println("Waiting goroutines")
	wg.Wait()
	
}