package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main(){
	var vName string = "Derek"
	var v1, v2 = 1.2, 1.3
	var v3 = "hello"
	v4 := 2.4
	v4 = 5.4

	fmt.Println(vName, v1, v2, v3, v4)
}

data types
var pl = fmt.Println
func main(){
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(25.5))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf(true))
}

to cast
var pl = fmt.Println
func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2) 

	cV3 := "500000000"
	cV4, err := strconv.Atoi(cV3)
	if err == nil {
		pl(cV4)
		pl(reflect.TypeOf(cV4))
	} else {
		pl(err)
	}
}

