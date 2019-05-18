package main

import "fmt"

func main() {
	// Longest way of declaring a variable
	// var name = "Besto"
	// var age = 53534895345
	const isSmart = true

	lovesGod := true

	// Average way of declaring a variable
	// name := "Best"
	// age := 21

	// Shorter way of declaring a variable
	name, age := "Best", 21

	fmt.Println(name, age)
	fmt.Println(name, lovesGod)

	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", lovesGod)
	fmt.Printf("%T\n", isSmart)
}
