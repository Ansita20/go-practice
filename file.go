package main

import(
	"fmt"
	"os"
	"log"
)

var pl = fmt.Println
func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
}