package main

import "fmt"

func log(text ...string) {
	fmt.Println(text[0])
}

func main() {
	log("Hello world")
}
