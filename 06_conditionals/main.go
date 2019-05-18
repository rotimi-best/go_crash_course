package main

import "fmt"

func log(text ...string) {
	fmt.Println(text[0])
}

func main() {
	x := 45
	y := 20

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// Else if
	color := "daf"

	if color == "red" {
		log("color is red")
	} else if color == "blue" {
		log("color is blue")
	} else {
		log("color is not blue nor red")
	}

	// Switch
	switch color {
	case "red":
		log("color is red")
	case "blue":
		log("color is blue")
	default:
		log("color is not blue nor red")
	}
}
