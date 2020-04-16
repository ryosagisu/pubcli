package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/urfave/cli"
)

var (
	EMPTY_TOPIC        = fmt.Errorf("TOPIC: Must be specified.")
	EMPTY_SUBSCRIPTION = fmt.Errorf("SUBSCRIPTION: Must be specified.")
)

func main() {
	ctx := context.Background()
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	client, err := pubsub.NewClient(ctx, "foobar")
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "topics",
			Aliases: []string{"t"},
			Usage:   "Manage Cloud Pub/Sub topics.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Creates one or more Cloud Pub/Sub topics.",
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_TOPIC
						}

						// gcloud pubsub topics create TOPIC [TOPIC ...]
						return CreateTopics(client, c.Args())
					},
				},
				{
					Name:  "delete",
					Usage: "Deletes one or more Cloud Pub/Sub topics.",
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_TOPIC
						}

						// gcloud pubsub topics delete TOPIC [TOPIC ...]
						return DeleteTopics(client, c.Args())
					},
				},
				{
					Name:  "list",
					Usage: "Lists Cloud Pub/Sub topics",
					Action: func(c *cli.Context) error {
						// gcloud pubsub topics list
						return ListTopics(client)
					},
				},
				{
					Name:  "list-subscriptions",
					Usage: "Lists Cloud Pub/Sub subscriptions from a given topic.",
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_TOPIC
						}

						// gcloud pubsub topics list-subscriptions TOPIC
						return ListSubscriptionsInTopic(client.Topic(c.Args().Get(0)))
					},
				},
				{
					Name:  "publish",
					Usage: "Publishes a message to the specified topic name",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "message",
							Value:    "MESSAGE",
							Usage:    "The body of the message to publish to the given topic name",
							Required: true,
						},
					},
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_TOPIC
						}

						// gcloud pubsub topics publish TOPIC [--message=MESSAGE]
						return Publish(client.Topic(c.Args().Get(0)), c.String("message"))
					},
				},
			},
		},
		{
			Name:    "subscriptions",
			Aliases: []string{"s"},
			Usage:   "Manage Cloud Pub/Sub subscriptions.",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Creates one or more Cloud Pub/Sub subscriptions.",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "topic",
							Value:    "TOPIC",
							Usage:    "topic name",
							Required: true,
						},
					},
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_SUBSCRIPTION
						}

						// gcloud pubsub subscriptions create SUBSCRIPTION [SUBSCRIPTION ...] --topic=TOPIC
						return CreateSubscription(client, c.String("topic"), c.Args())
					},
				},
				{
					Name:  "delete",
					Usage: "Deletes one or more Cloud Pub/Sub subscriptions.",
					Action: func(c *cli.Context) error {
						if c.NArg() == 0 {
							return EMPTY_TOPIC
						}

						// gcloud pubsub subscriptions delete SUBSCRIPTION [SUBSCRIPTION ...]
						return DeleteSubscription(client, c.Args())
					},
				},
				{
					Name:  "list",
					Usage: "Lists Cloud Pub/Sub subscriptions.",
					Action: func(c *cli.Context) error {
						// gcloud pubsub subscriptions list
						return ListSubscription(client)
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
