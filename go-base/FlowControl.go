package main

import (
	"fmt"
)

func main() {

	defer myfunc(1, 2, 3, 4, 5, 6, 7, 9, 6, 4)
	//if
	i := 13
	if i > 10 {
		fmt.Println("i is bigger than 10")
	} else {
		fmt.Println("i is small than 10")
	}

	//for
	sum := 0;
	for ; sum < 100; sum++ {
		sum += sum;
	}
	fmt.Println("the sum is:", sum)

	//switch
	flag := 6
	switch flag {
	case 4:
		fmt.Println("it is equal 4") //自带break,要想取消break=>使用fallthrough
	case 5:
		fmt.Println("it is equal 5")
	case 6:
		fmt.Println("it is equal 6")

	}

	//function
	x := 3
	y := 5
	xPLUSy, xTIMESy := sumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d + %d = %d\n", x, y, xTIMESy)

	//pointer
	xx := 3
	xx1 := add(3)
	fmt.Println("xx + 1 = ", xx1)
	fmt.Println("xx = ", xx)

	xxx := 3
	xxx1 := addWithPoint(&xxx)
	fmt.Println("xxx + 1 = ", xxx1)
	fmt.Println("xxx = ", xxx)

	//defer
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd elements of this slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("even elements of slice are:", even)

}

func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func myfunc(arg ... int) {

	for ii := 0; ii < len(arg); ii++ {
		println(arg[ii])
	}
}

func add(a int) int {
	a = a + 1
	return a
}

func addWithPoint(a *int) int {
	*a = *a + 1
	return *a
}

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

type testInt func(int) bool

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
