package main

import (
	"fmt"
	"math"

	"github.com/rotimi-best/go_crash_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(24))
	fmt.Println(math.Ceil(2.7))

	str, reveresedStr := "Hello", ""

	reveresedStr = strutil.Reverse(str)

	fmt.Println(str, reveresedStr)
}
