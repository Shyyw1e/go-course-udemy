package gpt_tasks

import (
	"context"
	"fmt"
)

func printUser(ctx context.Context) {
	v := ctx.Value("userID")
	
	userID, ok := v.(int)
	if !ok {
		fmt.Println("userID not found")
		return
	}
	fmt.Println("userID: ", userID)
}


func Task3() {
	ctx := context.WithValue(context.Background(), "userID", 123)
	printUser(ctx)
}