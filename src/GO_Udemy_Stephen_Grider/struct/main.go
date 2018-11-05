package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	dirx := person{
		firstName: "Dirx",
		lastName:  "Horneal",
		contact: contactInfo{
			email:   "Dirx@dirx.com",
			zipCode: 69699,
		},
	}
	//dirxPOinter := &dirx
	dirx.updateName("horneal")
	dirx.print()
}
func (pointerToPerson *person) updateName(newFirsName string) {
	(*pointerToPerson).firstName = newFirsName
}

func (p person) print() {
	fmt.Printf("%+v", p)

}
