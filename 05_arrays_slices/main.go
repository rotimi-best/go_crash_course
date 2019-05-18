package main

import "fmt"

func main() {
	// Array
	// var bookArr [2]string

	// Assigning values
	// bookArr[0] = "Mathew"
	// bookArr[1] = "Mark"

	// Declare and assign
	// bookArr := [2]string{"Matthew", "Mark"}

	// fmt.Println(bookArr)

	bookSlice := []string{"Matthew", "Mark", "Luke"}
	fmt.Println(len(bookSlice))
	fmt.Println(bookSlice[0:2])
}
