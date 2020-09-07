package main

import (
	"fmt"
	"github.com/acsauk/test-with-go/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		msg := fmt.Sprintf("FAIL: wanted 11, but recevied %d", sum)
		panic(msg)
	}

	add := math.Add(5, 10)
	if add != 15 {
		msg := fmt.Sprintf("FAIL: wanted 15, but recevied %d", add)
		panic(msg)
	}

	fmt.Println("PASS")
}