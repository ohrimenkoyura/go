package main

import "fmt"

func between(from, to int) []int {
	var result []int

	if from > to {
		result = []int{}
	} else {
		for i := from; i < to; i++ {
			result = append(result, i)
		}
	}
}

func main() {
	for i := range between(from:0, to:10) {
		switch i % 5 {
		case 1:
			fmt.Println(a:"fizz")
		case 2:
			fmt.Println(a:"bazz")
		case 3:
			fmt.Println(a:"gizz")
			fallthrough
		default:
			fmt.Println(a:"fizzbazz")
		}
	}
}
