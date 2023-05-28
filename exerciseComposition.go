package main

import "fmt"


type Person struct {
	name string
	age int
}

type Employee struct {
	p Person
	empId int
}

func (p *Person) printName () string {
	return p.name
}

func main() {
	e := Employee{
		p: Person{name: "Jane", age: 45},
		empId: 2332,
	}

	sayHello(e)
}

func sayHello (e Employee) string {
	fmt.Println(e.p.printName())
	return e.p.printName()
}