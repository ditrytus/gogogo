package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, callback := context.WithCancel(context.Background())
	ctx, _ := context.WithTimeout(context.Background(), 2 * time.Second)

	for i := 0; i < 5; i++ {
		go func(ctx context.Context, n int) {
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Println(n, ctx.Err())
		}(ctx, i)

		//if i == 4 {
		//	callback()
		//	break
		//}
	}

	<-ctx.Done()

	fmt.Println("Mamma mia")
}
