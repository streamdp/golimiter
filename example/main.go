package main

import (
	"context"
	"log"
	"time"

	"github.com/streamdp/golimiter"
)

func main() {
	var (
		allowed int

		now = time.Now()
	)

	ctx := context.Background()
	l := golimiter.New(ctx, time.Minute)

	for time.Since(now) <= 2*time.Second {
		if l.Allow(ctx, "key1", 10, time.Second) {
			allowed++
		}
	}

	log.Println("hits count:", allowed)
	log.Println("test passed: ", allowed == 20)
}
