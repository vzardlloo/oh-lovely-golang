package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0;
	for _, v := range a {
		total += v
	}
	//把值发送到channel
	c <- total
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y

	}
	close(c)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	//定义一个int类型的channel
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// fibonacci
	fi := make(chan int, 10)
	go fibonacci(cap(fi), fi)
	for i := range fi {
		fmt.Print(i)
		fmt.Print(" ")
	}

}
