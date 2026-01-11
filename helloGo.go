package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
var pl = fmt.Println
func main() {
	pl("Hello, Go!")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("hello", name)
	}else{
		log.Fatal(err)
	}
}