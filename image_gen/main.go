package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hibiken/asynq"
)

type ImageGenPayload struct {
	UserName string
}

const redisAddr = "127.0.0.1:6379"

// A list of task types.
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

func HandleImageGen(ctx context.Context, t *asynq.Task) error {
	var p ImageGenPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Creating image: user_id=%s", p.UserName)
	// Email delivery code ...
	return nil
}

// This should be a perma listener to send emails
func main() {
	fmt.Println("Starting listener 2")
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeImageResize, HandleImageGen)
	log.Fatal(server.Run(mux))
}
