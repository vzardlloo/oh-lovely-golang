package main

import "fmt"

func main() {

	ci := make(chan int)
	go func() {
		ci <- 3
	}()
	v := <-ci

	fmt.Println(v)

}
