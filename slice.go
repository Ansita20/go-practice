package main

import (
	"fmt"
)

var pl = fmt.Println;
func main(){
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"
	pl("Slice Size :", len(sl1))
	for i := 0 ; i < len(sl1) ; i++ {
		pl(sl1[i])
	}
	
	for _, v := range(sl1) {
		pl(v)
	}

	sArr := [5]int {1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl(sl3)
	pl("first two: ", sArr[:2])
	pl("Last two: ", sArr[3:])
	sArr[0] = 10;
	pl(sl3)
	sArr[0] = 1
	sl3 = append(sl3,12)
	pl(sl3)

	sl4 := make([]string, 6)
	pl(sl4)
	pl("sl4[0] : ", sl4[0])
}