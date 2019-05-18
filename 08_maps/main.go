package main

import "fmt"

func log(text ...map[string]string) {
	for i := 0; i < len(text); i++ {
		fmt.Println(text[i])
	}
}

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key
	// emails["Brad"] = "bradtraversty@go.com"
	// emails["Rob"] = "roboike@go.com"
	// emails["Larry"] = "larry@go.com"

	// Declare map and add key values
	emails := map[string]string{
		"Rob":   "roboike@go.com",
		"Larry": "larry@go.com",
	}

	log(emails)

	// Delete from map
	delete(emails, "Brad")
	log(emails)
}
