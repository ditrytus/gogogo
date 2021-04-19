package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	fmt.Println(ctx.Err())
	cancel()
	fmt.Println(ctx.Err())
}
