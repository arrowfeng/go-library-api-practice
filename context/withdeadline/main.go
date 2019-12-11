package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	d := time.Now().Add(time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("oversleep")
	case <-ctx.Done():
		fmt.Println("cancel:", ctx.Err())
	}
}
