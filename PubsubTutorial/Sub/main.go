package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	projectID := "aftership-dev"

	mu1 := sync.Mutex{}
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println("create client error:", err)
	}

	mu2 := sync.Mutex{}
	sub1Received := 0
	sub2Received := 0

	for i := 0; i < 2; i++ {
		wg1.Add(1)
		go sub1Consume("sub1", &sub1Received, &mu1, ctx, client, "sub1-processor"+strconv.Itoa(i+1), &wg1)
	}

	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go sub1Consume("sub2", &sub2Received, &mu2, ctx, client, "sub2-processor"+strconv.Itoa(i+1), &wg2)
	}

	wg1.Wait()
	wg2.Wait()
}

func sub1Consume(subID string, sub1Count *int, mu *sync.Mutex, ctx context.Context, client *pubsub.Client, proceessName string, wg *sync.WaitGroup) {
	sub := client.Subscription(subID)
	cctx, cancelFunc := context.WithCancel(ctx)
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(proceessName+"receive msg:", string(msg.Data))
		msg.Ack()
		*sub1Count++
		if *sub1Count == 10 {
			fmt.Println("receive 10,cancel")
			cancelFunc()
		}
	})
	if err != nil {
		fmt.Println("err occur")
	}
	wg.Done()
}
