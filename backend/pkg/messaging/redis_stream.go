package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const AnalyticsStream = "stream:analytics"
const AnalyticsGroup  = "group:analytics"

type AnalyticsPayload struct {
	LinkID    int64     `json:"link_id"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"ua"`
	Referrer  string    `json:"ref"`
	Timestamp time.Time `json:"ts"`
}

type StreamManager struct {
	client *redis.Client
}

func NewStreamManager(client *redis.Client) *StreamManager {
	return &StreamManager{client: client}
}

// Publish adds a message to the stream
func (s *StreamManager) Publish(ctx context.Context, stream string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshalling payload: %w", err)
	}

	return s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: map[string]interface{}{"data": data},
	}).Err()
}

// EnsureGroup creates the consumer group if it doesn't exist
func (s *StreamManager) EnsureGroup(ctx context.Context, stream, group string) {
	s.client.XGroupCreateMkStream(ctx, stream, group, "$")
}

// Consume reads messages from the stream for a group
func (s *StreamManager) Consume(ctx context.Context, stream, group, consumer string, handler func(data []byte) error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msgs, err := s.client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    group,
				Consumer: consumer,
				Streams:  []string{stream, ">"},
				Count:    10,
				Block:    5 * time.Second,
			}).Result()

			if err != nil {
				if err != redis.Nil {
					fmt.Printf("Error reading from stream: %v\n", err)
				}
				continue
			}

			for _, xStream := range msgs {
				for _, msg := range xStream.Messages {
					data := msg.Values["data"].(string)
					if err := handler([]byte(data)); err == nil {
						s.client.XAck(ctx, stream, group, msg.ID)
					} else {
						fmt.Printf("Error handling message %s: %v\n", msg.ID, err)
					}
				}
			}
		}
	}
}
