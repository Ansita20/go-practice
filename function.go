package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("hello world!")
}

func main() {
	sayHello()
	sum(1 ,6)
	pl(quitient(3, 0))
	pl(getSum2(2,3,4,5))

	arr := []int{1,2,3,4} 
	pl(getArray(arr))
}

func sum(x int ,y int)string{
	x = x + 2
	pl(x)
	x = x + y
	pl(x)

	return "hello there!"
}

func quitient(x int , y int)(int, error){
	if y == 0 {
		return 0, fmt.Errorf("division by 0")
	}
	return x/y, nil
}

func getSum2 (nums ...int)int {
	sum := 0
	for _, num := range(nums){
		sum += num
	}
	return sum
}

func getArray(arr []int) int{
	sum := 0
	for _, num := range arr{
		sum += num
	}
	return sum
}