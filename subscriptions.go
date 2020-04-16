package main

// [START pubsub_create_pull_subscription]
import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

func CreateSubscription(client *pubsub.Client, topicID string, subID []string) error {
	ctx := context.Background()

	for _, id := range subID {
		sub, err := client.CreateSubscription(ctx, id, pubsub.SubscriptionConfig{
			Topic:       client.Topic(topicID),
			AckDeadline: 20 * time.Second,
		})
		if err != nil {
			return fmt.Errorf("CreateSubscription: %v", err)
		}
		fmt.Printf("Created subscription: %v\n", sub)
	}
	return nil
}

// [END pubsub_create_pull_subscription]

// [START pubsub_delete_subscription]

func DeleteSubscription(client *pubsub.Client, subID []string) error {
	ctx := context.Background()

	for _, id := range subID {
		sub := client.Subscription(id)
		if err := sub.Delete(ctx); err != nil {
			return fmt.Errorf("Delete: %v", err)
		}
		fmt.Printf("Subscription %q deleted.", id)
	}
	return nil
}

// [END pubsub_delete_subscription]

// [START pubsub_list_subscriptions]

func ListSubscription(client *pubsub.Client) error {
	ctx := context.Background()

	it := client.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Next: %v", err)
		}
		fmt.Printf("%v\n", s)
	}
	return nil
}

// [END pubsub_list_subscriptions]
