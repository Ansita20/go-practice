package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

// var pl = fmt.Println

// func main() {
// 	for x := 5 ; x >= 1 ; x-- {
// 		pl(x)
// 	}
// }

// for used as while
// var pl = fmt.Println
// func main() {
// 	fx := 0
// 	for fx < 6 {
// 		pl(fx)
// 		fx++;
// 	}
// }

//infinite loop
//while true loop
var pl = fmt.Println
func main(){
	seed := time.Now().Unix()
	rand.Seed(seed)
	randNum := rand.Intn(50) + 1;
	for true {
		pl(randNum)
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if(err != nil){
			log.Fatal(err)
		}

		if iGuess > randNum {
			pl("Too high!")
		}else if iGuess < randNum {
			pl("Too Low!")
		}else {
			pl("You guessed it!")
			break
		}
	}
}