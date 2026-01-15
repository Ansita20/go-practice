package main

import(
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

var pl = fmt.Println
func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2,3,5,7,11}
	var PrimeArr []string
	for _, i := range iPrimeArr {
		PrimeArr = append(PrimeArr, strconv.Itoa(i))
	}
	for _, num := range PrimeArr{
		_,err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime : ", scan1.Text())
	}
	if err := scan1.Err() ; err != nil{
		log.Fatal(err)
	}
}