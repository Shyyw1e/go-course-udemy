package gpt_tasks

import (
	"context"
	"fmt"
)

func logRequestInfo(ctx context.Context) {
	requestID, ok1 := ctx.Value("requestID").(string)
	userRole, ok2 := ctx.Value("userRole").(string)
	traceID, ok3 := ctx.Value("traceID").(string)
	
	if !ok1 || !ok2 || !ok3 {
		fmt.Println("Some values are absent")
		return
	} 

	fmt.Println("Logging request info:")
	fmt.Println("Request ID: ", requestID)
	fmt.Println("User Role : ", userRole)
	fmt.Println("Trace ID  : ", traceID)
}

func processBusinessLogic(ctx context.Context) {
	fmt.Println("\nProcessing business logic ")
	logRequestInfo(ctx)
}

func handleRequest(ctx context.Context) {
	fmt.Println("Handling request . . .")
	processBusinessLogic(ctx)
}

func Task3_2() {
	ctx := context.WithValue(context.Background(), "requestID", "req-1234")
	ctx = context.WithValue(ctx, "userRole", "admin")
	ctx = context.WithValue(ctx, "traceID", "trace-abcde")

	handleRequest(ctx)
}