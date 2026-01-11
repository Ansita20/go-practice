package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func main() {
	arr1 := []int {1,2,3}
	for _, val := range arr1{
		pl(val)
	}

	var arr4 [5]int

	arr2 := [5]int{1,2,3,4,5}
	pl("0 index : ",arr2[0])
	pl("length: ", len(arr2))

	reader := bufio.NewReader(os.Stdin)

	for i := 0 ; i < len(arr4) ; i++{
		fmt.Fscan(reader,&arr4[i])
	}
	pl(arr4)

	aStr1 := "abcde"
	rArr := []rune(aStr1)

	for _, v := range rArr{
		pl("Rune Array : ", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("Byte to String: ", bStr)
}