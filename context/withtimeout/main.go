package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func work(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Doing", i)
		case <-ctx.Done():
			fmt.Println("Cancel", i)
			return ctx.Err()

		}
	}
	return nil
}

func main() {

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("start work")

	wg.Add(1)
	go work(ctx, &wg)
	wg.Wait()

	fmt.Println("finished")

}
