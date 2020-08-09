package main

import "fmt"

type person struct {
	lastName  string
	firstName string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	mars := person{
		firstName: "Mars",
		lastName:  "CHEN",
		contactInfo: contactInfo{
			email:   "mars@gmail.com",
			zipCode: 94000,
		},
	}

	// &variable 位置: 可拿到這個變量指向的值的內存地
	// pointerMars := &mars
	mars.updateName("Sally")
	mars.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *pointer 取值: 拿到這個內存地址指向的值
func (pointerToPerson *person) updateName(newFirtsName string) {
	(*pointerToPerson).firstName = newFirtsName
}
