package main

import (
	"fmt"
)

var pl = fmt.Println

func strpr (myptr *int) int{
	*myptr = 12
	return *myptr
}

func arrraP(arr *[4]int){
	for x := 0 ; x < 4 ; x++{
		arr[x] *= 2
	}
}

func floati(arr ...float64) float64 {
	var sum float64 = 0.0
	var numsize float64 = float64(len(arr))
	for _, val := range arr {
		sum += val
	}
	return (sum / numsize)
}

func main() {
	f4 := 10
	var f42 *int = &f4
	pl(f42)
	pl(*f42)
	f4 = 12
	pl(f42)

	f3 := 10
	pl(f3)
	strpr(&f3)
	pl(f3)

	pArr := [4]int {1,2,3,4}
	arrraP(&pArr)
	pl(pArr)

	parr := []float64{1,2,3,4}
	pl("%.3f\n",floati(parr...))
}

