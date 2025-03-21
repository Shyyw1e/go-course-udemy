package concurency

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)

	parse(ctx)
}

func parse(ctx context.Context) {
	for {
		select {
		case <-time.After(time.Second * 2):
			fmt.Println("parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceeded")
			return
		}
	}
}
