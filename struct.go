package main

import(
	"fmt"
)

var pl = fmt.Println

// type myConstraint interface {
// 	int | float64
// }

// func getSum[T myConstraint] (x T, y T) T {
// 	return x + y;
// }

// func main(){
// 	pl(getSum(2.0,3.3))
// }

// type customer struct {
// 	name string
// 	address string
// 	bal float64
// }

// func getCustInfo(c customer){
// 	pl("%s owes us %.2f\n", c.name, c.bal)
// }

// func newCustAdd(c *customer, address string){
// 	c.address = address
// }

// func main() {
// 	var tS customer
// 	tS.name = "Tom Smith"
// 	tS.address = "London Street 54"
// 	tS.bal = 22.34
// 	pl(tS)
// 	getCustInfo(tS)
// 	newCustAdd(&tS, "123 South st")
// 	pl("Address : ", tS.address)
// 	sS := customer{"Sam Smith", "123 lane", 0.0}
// 	pl("Name : ", sS.name)
// }

//type provides a data type to the var 
// type rectangle struct {
// 	length, height float64
// }

// func (r rectangle) Area() float64{
// 	return r.length * r.height
// }

// func main(){
// 	rect1 := rectangle{10.0, 15.0}
// 	pl("Rect Area : ", rect1.Area())
// }

type contact struct{
	fName string
	lName string
	phone string
}

type buisness struct{
	name string
	address string
	contact 
}

func (b buisness) info() {
	pl("Contact at",b.name,"is",b.contact.fName, b.contact.lName)
}

func main() {	
	bus1 := buisness{
		name:    "ABC Plumbing",
		address: "234 North St",
		contact: contact{
			fName: "Robert",
			lName: "Downey",
			phone: "17653498",
		},
	}

	bus1.info()
}