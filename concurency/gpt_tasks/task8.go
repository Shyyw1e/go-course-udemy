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
	reqID, ok := ctx.Value("requestID").(string)
	userRole, ok1 := ctx.Value("userRole").(string)
	defer wg.Done()
	if !ok {
		fmt.Println("Missing requestID in context")
	}
	if !ok1 {
		fmt.Println("Missing userRole in context")
	}
	if !ok || !ok1 {
		fmt.Println("❌ Missing requestID or userRole in context")
		return
	}

	fmt.Printf("Request ID:\t%s\nUser role:\t%s\n", reqID, userRole)
	
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


func Task8() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "requestID", "req-789")
	ctx = context.WithValue(ctx, "userRole", "admin")
	wg := &sync.WaitGroup{}
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	defer cancel()
	
	go func() {
		<-signalCh
		fmt.Println("Shutting down gracefully")
		cancel()
	}()
	wg.Add(3)
	go worker2(ctx, wg, 1)
	go worker2(ctx, wg, 2)
	go worker2(ctx, wg, 3)

	<-ctx.Done()
	
	fmt.Println("waiting for goroutines")
	wg.Wait()
	fmt.Println("Goroutines completed")
	

}