package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Context Tutorial")
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("oh no, deadline exceeded")
	}

	time.Sleep(2 * time.Second)
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timed out")
			return
		default:
			fmt.Println("do something")
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println(rID)
}

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}
