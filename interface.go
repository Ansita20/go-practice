package main

import(
	"fmt"
)

var pl = fmt.Println

type Animal interface{
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	pl("cat Attacks its Prey")
}

func (c Cat) Name() string{
	return string(c)
}

func (c Cat) AngrySound(){
	pl("cat says Hisssss")
}

func (c Cat) HappySound(){
	pl("cat says Purrrrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()
	kitty.HappySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}