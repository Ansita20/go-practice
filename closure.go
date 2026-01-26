package main

import (
	"fmt"
)

var pl = fmt.Println

func useFunc(f func(int, int) int, x, y int){
	pl("Answer :", (f(x,y)))
}

func sumvals(x, y int) int {
	return x + y
}

func main() {
	intSum := func(x, y int) int {return x + y}
	pl("5 + 4 =", intSum(5,4))

	sample1 := 1
	changeVar := func() {
		sample1 += 1
	}
	changeVar()
	pl("Sample1 =", sample1)
	useFunc(sumvals, 5 ,8)
}