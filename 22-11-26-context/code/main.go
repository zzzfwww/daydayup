package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	c1, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	c2 := context.WithValue(c1, "key", "value")
	c3 := context.WithValue(c2, "key2", "value3")
	// c4, cancelFunc := context.WithTimeout(c3, 100*time.Millisecond)
	c4, cancelFunc := context.WithTimeout(c3, 100*time.Second)
	defer cancelFunc()
	deadline, ok := c4.Deadline()
	fmt.Println(deadline, ok)
}
