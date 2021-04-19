package main

import "fmt"

func main() {
	c := make(chan struct{foo int})
	close(c)
	x := <-c
	fmt.Println("done", x)
}
