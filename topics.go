package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

// [START pubsub_create_topic]

func CreateTopics(client *pubsub.Client, topicID []string) error {
	ctx := context.Background()

	for _, id := range topicID {
		t, err := client.CreateTopic(ctx, id)
		if err != nil {
			return fmt.Errorf("CreateTopic: %v", err)
		}
		fmt.Printf("Topic created: %v\n", t)
	}
	return nil
}

// [END pubsub_create_topic]

// [START pubsub_delete_topic]

func DeleteTopics(client *pubsub.Client, topicID []string) error {
	ctx := context.Background()

	for _, id := range topicID {
		t := client.Topic(id)
		if err := t.Delete(ctx); err != nil {
			return fmt.Errorf("Delete: %v", err)
		}
		fmt.Printf("Deleted topic: %v\n", t)
	}
	return nil
}

// [END pubsub_delete_topic]

// [START pubsub_list_topics]

func ListTopics(client *pubsub.Client) error {
	ctx := context.Background()

	it := client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next: %v", err)
		}
		fmt.Printf("%v\n", topic)
	}
	return nil
}

// [END pubsub_list_topics]

// [START pubsub_list_topic_subscriptions]

func ListSubscriptionsInTopic(topic *pubsub.Topic) error {
	ctx := context.Background()

	it := topic.Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next: %v", err)
		}
		fmt.Printf("%v\n", sub)
	}
	return nil
}

// [END pubsub_list_topic_subscriptions]

// [START pubsub_publish]

func Publish(topic *pubsub.Topic, msg string) error {
	ctx := context.Background()

	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

// [END pubsub_publish]
