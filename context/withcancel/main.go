package main

import (
	"context"
	"fmt"
	"time"
)

func Stream(ctx context.Context, out chan<- string) {

	for {

		name := "halo"
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("context is canceled:", ctx.Err())
			return
		case out <- name:
		}
	}

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan string)

	go Stream(ctx, out)

	count := 0
	for name := range out {
		fmt.Println(name)
		time.Sleep(time.Second)
		if count == 2 {
			break
		}
		count++
	}

	cancel()
	time.Sleep(time.Second * 1)

}
