package main

import "fmt"

func greeting(text string, name string) string {
	return text + " " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	const welcome = "Welcome to Go"
	var greetingText = greeting(welcome, "Best")

	fmt.Println(greetingText)

	var sum = getSum(1, 2)
	fmt.Println(sum)
}
