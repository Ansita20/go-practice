package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println
func main() {
	pl(os.Args)
	args := os.Args[1:]
	var iargs = []int {}
	for _, i := range args{
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iargs = append(iargs, val)
	}
	max := 0
	for _, val := range iargs {
		if val > max {
			max = val
		}
	}
	pl(max)
}
