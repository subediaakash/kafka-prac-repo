package events

import "time"

type PaymentEvents struct {
	EventId   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PaymentId string    `json:"payment_id"`
	PayerId   string    `json:"payer_id"`
	PayerType string    `json:"payer_type"`
	Purpose   string    `json:"purpose"`
	Currency  string    `json:"currency"`
	OccuredAt time.Time `json:"occured_at"`
}
