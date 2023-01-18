package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)
	sleepAndTalk(ctx, 5*time.Second, "Hello")
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
