package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

// var pl = fmt.Println
// func main() {
// 	f, err := os.Create("file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	arrp := []int {1,2,3,4,5}
// 	var iarrp []string
// 	for _, i := range arrp {
// 		iarrp = append(iarrp, strconv.Itoa(i))
// 	}

// 	for _, num := range iarrp {
// 		_, err := f.WriteString(num + "\n")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	f,err = os.Open("file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	scan1 := bufio.NewScanner(f)
// 	for scan1.Scan() {
// 		pl("Prime: ", scan1.Text())
// 	}
// }

var pl = fmt.Println

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	arrp := []int {1,2,3,4,5}
	var iarrp []string
	for _, num := range arrp {
		iarrp = append(iarrp, strconv.Itoa(num))
	}

	for _,num := range iarrp {
		_, err := os.
	}
}