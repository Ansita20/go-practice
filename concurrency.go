package main

import (
	"fmt"
	// "time"
)

var pl = fmt.Println

// func printTo15(){
// 	for i := 1; i <= 15; i++ {
// 		pl("Fun 1 :", i)
// 	}
// }

// func printTo10(){
// 	for i := 1; i <= 10; i++ {
// 		pl("Fun 2 :", i)
// 	}
// }

// func main() {
// 	go printTo15()
// 	go printTo10()
// 	time.Sleep(7 * time.Second)
// }

func nums1(channel chan int){
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int){
	channel <- 4
	channel <- 5
	channel <- 6
}

func main(){
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<- channel1)
	pl(<- channel1)
	pl(<- channel1)
	pl(<- channel2)
	pl(<- channel2)
	pl(<- channel2)

}