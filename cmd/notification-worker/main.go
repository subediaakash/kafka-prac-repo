package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/subediaakash/eventflow/internal/events"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	group := os.Getenv("GROUP")
	if group == "" {
		group = "notification-service"
	}
	log.Printf("joining consumer group %q", group)

	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup(group),
		kgo.ConsumeTopics("application-events-v2"),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	ctx := context.Background()
	log.Println("Notification worker started , waiting for the messages ....")

	for {
		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Printf("fetch errors: %v", errs)
			continue
		}
		fetches.EachRecord(func(r *kgo.Record) {
			var event events.ApplicationEvent

			if err := json.Unmarshal(r.Value, &event); err != nil {
				log.Printf("skipping bad messge partition = %dd offset = %d : %v", r.Partition, r.Offset, err)
				return
			}

			log.Printf("got event type=%s application=%s candidate=%s",
				event.EventType, event.ApplicationID, event.CandidateID)
		})
	}

}
