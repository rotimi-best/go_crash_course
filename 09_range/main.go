package main

import "fmt"

func log(text ...string) {
	fmt.Println(text[0])
}

func main() {
	ids := []int{23, 445, 6562, 2, 367, 7, 0}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Loop through ids not using index
	for _, id := range ids {
		fmt.Printf("%d - ID\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range with maps
	emails := map[string]string{
		"Rob":   "roboike@go.com",
		"Larry": "larry@go.com",
	}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		log("Name: " + k)
	}
}
