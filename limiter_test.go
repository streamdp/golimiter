package golimiter

import (
	"context"
	"testing"
	"time"
)

func TestLimiter(t *testing.T) {
	type args struct {
		key    string
		rate   int
		period time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "rate 100 per second",
			args: args{
				key:    "key1",
				rate:   100,
				period: time.Second,
			},
		},
		{
			name: "rate 10 per second",
			args: args{
				key:    "key1",
				rate:   10,
				period: time.Second,
			},
		},
		{
			name: "rate 150 per second",
			args: args{
				key:    "key1",
				rate:   150,
				period: 100 * time.Millisecond,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(context.Background())
			allowed := 0

			now := time.Now()
			for time.Since(now) <= tt.args.period {
				if l.Allow(tt.args.key, tt.args.rate, tt.args.period) {
					allowed++
				}
			}

			if allowed != tt.args.rate {
				t.Errorf("rate = %v, want %v", allowed, tt.args.rate)
			}
		})
	}
}
