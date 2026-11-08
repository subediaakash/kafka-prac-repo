package events

import "time"

type ApplicationEvent struct {
	EventID       string    `json:"event_id"`
	EventType     string    `json:"event_type"`
	ApplicationID string    `json:"application_id"`
	CandidateID   string    `json:"candidate_id"`
	OccurredAt    time.Time `json:"occurred_at"`
}
