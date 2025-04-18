## Golimiter

Library to limiting something: https request, loops, sends to the channel & etc.
Written on **golang** and used [microcache](https://github.com/streamdp/microcache) library.

### Usage
Example library usage for limiting **doSomething()** function on 10 calls per second:
```go
package main

import (
    "context"
    "time"
    "github.com/streamdp/golimiter"
)

func main() {
    ctx := context.Background()
    l := golimiter.New(ctx, time.Minute)
    if l.Allow(ctx, "key1", 10, time.Second) {
        doSomething()
    }
}
```
