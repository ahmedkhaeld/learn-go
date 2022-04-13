package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

// mySleepAndTalk either it has enough time to fire a message or does not have time so context call
// Done() that close, because the timeout is elapsed and log the cancellation err
func mySleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}

}
