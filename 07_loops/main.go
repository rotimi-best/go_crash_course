package main

import "fmt"

func log(text ...string) {
	fmt.Println(text[0])
}

func main() {
	// Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Shorted method
	for i := 1; i <= 10; i++ {
		log("I will run 10 times")
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			log("FizzBuzz")
		} else if i%3 == 0 {
			log("Fizz")
		} else if i%5 == 0 {
			log("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
