package main

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
)

func main() {
	projectID := "aftership-test"
	subID := "ggq-connectors-automizely-orders-pull-test"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Consume 10 messages.
	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		panic(err)
	}
}
