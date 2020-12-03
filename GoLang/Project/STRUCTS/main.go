package main

import "fmt"

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

	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Naved", lastName: "Ahmed"}
	fmt.Println(alex)

	var aron person
	fmt.Println(aron)
	aron.firstName = "Aron"
	aron.lastName = "Ahmed"
	aron.contact.email = "test@123.com"
	aron.contact.zipCode = 95117
	fmt.Println(aron)
	fmt.Printf("%+v", aron)
	fmt.Println("--new line now--")
	aron.print()
	fmt.Println("--new line now--")
	//aron.updateFirstName("Test")
	aron.print()
	aronPointer := &aron
	aronPointer.updateName("jimmy")
	aron.print()
	fmt.Println("--new line now--")
	aron.updateName("timmy")
	aron.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

/*func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}*/
func (p person) print() {
	fmt.Printf("%+v", p)
}
