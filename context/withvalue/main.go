package main

import (
	"context"
	"fmt"
	"time"
)

var (
	key = "halo"
)

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "is canceled")
			return
		default:
			fmt.Println(ctx.Value(key), "normal print")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx, key, "zdf")

	go watch(valueCtx)

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(time.Second)

}
