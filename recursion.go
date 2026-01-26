package main

import (
	"fmt"
)

var pl = fmt.Println

func fact (a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a - 1)
}

func main() {
var a = fact(4)
pl(a)
}
