package main

import (
	"context"
	"fmt"
	"time"
)

func foo2(ctx context.Context, done chan struct{}) {
	fmt.Println("beginning work")
	<-ctx.Done()
	fmt.Println("finishing work")
	close(done)
}

func main() {
	backgroundContext := context.Background()
	//ctx, _ := context.WithDeadline(backgroundContext, time.Now().Add(1*time.Second))
	ctx, _ := context.WithTimeout(backgroundContext, 1*time.Second)
	done := make(chan struct{})
	done2 := make(chan struct{})
	done3 := make(chan struct{})
	go foo2(ctx, done)
	go foo2(ctx, done2)
	go foo2(ctx, done3)
	<-done
	<-done2
	<-done3
}
