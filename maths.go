package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

var pl = fmt.Println
func main() {
	pl("5 + 3 = ", 5+3)
	pl("5 - 3 = ", 5-3)
	pl("5 * 3 = ", 5*3)
	pl("5 / 3 = ", 5/3)

	mInt := 5
	mInt++
	pl(mInt)

	// 1/1/1970
	seedsecs := time.Now().Unix()
	pl("Unix Time: ", seedsecs)

	rand.Seed(seedsecs)
	randNum := rand.Intn(50) + 1
	pl("Random Number: ", randNum)

	seed := time.Now()
	pl(seed)

	pl("Pi value: ", math.Pi)
	pl("square root of 27: ",math.Sqrt(27))
	pl("log of 2: ",math.Log(2))
	pl("2 to the power 4 : ", math.Pow(2,4))
	pl("Absolute of 4.4: ", math.Abs(4.4))
	pl("Ceil of 4.4: ", math.Ceil(4.4))
	pl("Floor of 4.4: ", math.Floor(4.4))
	pl("Round of 4.4: ", math.Round(4.4))
	pl("Max of 4 and 9: ", math.Max(4,9))
	pl("Min of 4, 9: ", math.Min(4,9))
	pl("Sin of 4: ", math.Sin(4))

	r := 5
	pl("%v", r)
}