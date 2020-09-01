package main

import "fmt"

type person struct {
	lastName  string
	firstName string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.lastName, `says, "Good moring, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.lastName, sa.firstName, `says, "Shaken, not stirred.`)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	xi := []int{2, 4, 7, 9}
	fmt.Println(xi)

	m := map[int]string{
		3: "hello",
		4: "uyyo",
	}

	fmt.Println(m)

	mars := person{
		"CHEN",
		"Mars",
	}

	sal := secretAgent{
		person{
			"Sally",
			"CHEN",
		},
		true,
	}

	// fmt.Printf("%+v", mars)
	// mars.speak()
	// sal.person.speak()
	saySomething(mars)
	saySomething(sal)

}
