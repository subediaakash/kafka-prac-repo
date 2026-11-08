package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/subediaakash/eventflow/internal/events"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(brokers []string) (*Producer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
	)
	if err != nil {
		return nil, err
	}

	return &Producer{
		client: client,
	}, nil
}

func (p *Producer) PublishApplicationSubmitted(ctx context.Context, event events.ApplicationEvent) error {
	value, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("marshal event: %w", err)
	}

	record := &kgo.Record{
		Topic: "application-events-v2",
		Key:   []byte(event.ApplicationID),
		Value: value,
	}
	result := p.client.ProduceSync(ctx, record)
	if err := result.FirstErr(); err != nil {
		return fmt.Errorf("produce event: %w", err)
	}

	return nil

}

func (p *Producer) Close() {
	p.client.Close()
}
