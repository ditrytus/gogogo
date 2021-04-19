package main

import (
	"context"
	"fmt"
	"time"
)

func foo(ctx context.Context, done chan struct{}) {
	fmt.Println("beginning work")
	<- ctx.Done()
	fmt.Println("finishing work")
	close(done)
}

func main() {
	backgroundContext := context.Background()
	ctx, cancel := context.WithCancel(backgroundContext)
	ctx2, _ := context.WithCancel(ctx)
	done := make(chan struct{})
	done2 := make(chan struct{})
	go foo(ctx, done)
	go foo(ctx2, done2)
	time.Sleep(1*time.Second)
	cancel()
	<- done
	<- done2
}
