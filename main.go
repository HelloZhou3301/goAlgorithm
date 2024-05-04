package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 2, 1}
	limit := 3
	res := numRescueBoats(arr, limit)
	fmt.Print("res:", res)
}
