// structs2
// Make me compile!
//
package main

import "fmt"

type ContactDetails struct {
    string
}
type Person struct {
	// don't just create the phone field here. embed a new struct
    phone ContactDetails
	name string
	age int
}

func main() {
	contactDetails := ContactDetails{"+11 90 80 70"}
	person := Person{name: "John", age: 32, phone: contactDetails}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
