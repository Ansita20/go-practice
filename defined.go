package main

import (
	"fmt"
)

var pl = fmt.Println

type TSp float64
type TBs float64
type Ml float64

func tspToML (tsp TSp) Ml {
	return Ml(tsp * 4.92)
}

func tbsToML (tbs TBs) Ml {
	return Ml(tbs * 14.79)
}

func (tsp TSp) toTSp() Ml {
	return Ml(tsp * 4.92)
}

func (tbs TBs) toTBs() Ml {
	return Ml(tbs * 14.79)
}

func main() {
	r := tspToML(23)
	pl(r)

	p := tbsToML(23)
	fmt.Printf("%.2f\n", p)

	tsp1 := TSp(23)
	pl(tsp1.toTSp())
}