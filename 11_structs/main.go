package main

import (
	"fmt"
	"math"
	"strconv"
)

// Person Define person struct
type Person struct {
	name           string
	age            int
	favouriteColor string
}

// Greeting function (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.name + " and my age is " + strconv.Itoa(p.age)
}

// hasBirthday function (pointer reciever)
func (p *Person) moreMoney() {
	p.age = int(math.Pow(float64(p.age), 3))
}

func main() {
	// Long way
	person1 := Person{name: "Gbenga", age: 42, favouriteColor: "red"}

	// Short way
	person2 := Person{"Brad", 45, "black"}

	person2.age--

	fmt.Println(person1, person2)
	fmt.Println(person1.age, person2.favouriteColor)

	person2.moreMoney()
	fmt.Println(person2.greet())

}
