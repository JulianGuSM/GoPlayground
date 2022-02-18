package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
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
	id, rErr := result.Get(ctx)
	if rErr != nil {
		fmt.Println("publish err:\n", rErr)
	}
	fmt.Println("message id :", id)
}
