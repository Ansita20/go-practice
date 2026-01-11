package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println
func main() {
	now := time.Now()
	pl("Current Time: ", now.Year(), now.Day(), now.Month(), now.Hour(), now.Minute(), now.Second())
}
