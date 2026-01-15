package main

import(
	"fmt"
)

var pl = fmt.Println

type myCust struct{
	name string
	address int
	bal float64
}

func main() {
	var tS myCust
	tS.name = "ansita"
	pl(tS.name)
	tS.address = 12334
	pl(tS.address) 
}