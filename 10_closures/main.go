package main

import "fmt"

func log(text ...int) {
	fmt.Println(text[0])
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()

	for i := 0; i < 10; i++ {
		log(sum(i))
	}
}
