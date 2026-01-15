package main

import (
	"fmt"
)

var pl = fmt.Println
func main() {
	// var myMap map [keyType]valueType
	var heroes map[string]string
	heroes = make(map[string]string)

	// villains := make(map[string]string)
	heroes["Batman"] = "Bruce Wayne"

	for k,v := range heroes {
		fmt.Println("%s is %s \n", k , v)
	}

	superPets := map[int]string{1: "Krypto", 2: "Batman Hound"}
	for _, r := range superPets{
		pl(r);
	}
}