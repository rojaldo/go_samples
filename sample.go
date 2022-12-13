package main

import ("fmt" 
		"sample/person")



func main() {
	a := person.Person{}
	a.SetName("John")
	a.SetAge(20)
	z := person.Person{}
	z.SetName("Jane")
	fmt.Println(a, z)
	fmt.Println(a)

}