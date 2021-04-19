package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func foo3(ctx context.Context) error {
	fmt.Println("beginning work")
	<-ctx.Done()
	fmt.Println("finishing work")
	return nil
}

func main() {
	backgroundContext := context.Background()
	//ctx, _ := context.WithDeadline(backgroundContext, time.Now().Add(1*time.Second))
	ctx, _ := context.WithTimeout(backgroundContext, 1*time.Second)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := foo3(ctx)
			if err != nil {
				fmt.Println("stop immediately!")
				return
			}
			fmt.Println("success")
		}()
	}
	wg.Wait()
	fmt.Println("really done")
}
