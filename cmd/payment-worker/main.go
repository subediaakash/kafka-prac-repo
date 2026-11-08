package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/subediaakash/eventflow/internal/events"
	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup("payment-service"),
		kgo.ConsumeTopics("payment-topic"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	log.Println("payment worker started, waiting for messages...")

	for {
		fetches := client.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {
			log.Printf("fetch errors: %v", errs)
			continue
		}

		fetches.EachRecord(func(r *kgo.Record) {
			var event events.PaymentEvents

			if err := json.Unmarshal(r.Value, &event); err != nil {
				log.Printf("skipping bad message partition=%d offset=%d: %v",
					r.Partition, r.Offset, err)
				return
			}

			log.Printf("got payment event type=%s payment=%s payer=%s (%s) purpose=%s currency=%s",
				event.EventType, event.PaymentId, event.PayerId, event.PayerType,
				event.Purpose, event.Currency)
		})
	}
}
