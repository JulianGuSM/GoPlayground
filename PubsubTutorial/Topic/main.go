package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func main() {
	projectID := "aftership-dev"
	topicID := "topic-gsm"
	msg := "hello world"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println("create client error:\n", err)
	}

	t := client.Topic(topicID)
	for i := 0; i < 10; i++ {
		sendMsgToTopic(msg, t, ctx)
	}
}

func sendMsgToTopic(msg string, topic *pubsub.Topic, ctx context.Context) {
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	id, r_err := result.Get(ctx)
	if r_err != nil {
		fmt.Println("publish err:\n", r_err)
	}
	fmt.Println("message id :", id)
}
