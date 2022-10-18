package bq

import (
	"context"
	"log"
)

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// ProcessLogEntry processes a Pub/Sub message from Cloud Logging.
func ProcessLogEntry(ctx context.Context, m PubSubMessage) error {
	log.Printf("Log entry data: %s", string(m.Data))
	return nil
}
