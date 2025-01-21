package main

import (
	"context"
	"fmt"
	"time"

	"github.com/streamdp/golimiter"
)

func main() {
	var (
		allowed int

		now = time.Now()
	)

	l := golimiter.New(context.Background(), time.Minute)

	for time.Since(now) <= 2*time.Second {
		if l.Allow("key1", 10, time.Second) {
			allowed++
		}
	}

	fmt.Println("hits count:", allowed)
	fmt.Println("test passed: ", allowed == 20)
}
