package main

import "fmt"

func main() {

	intSlices := []int{1, 2, 4, 6, 8, 12}

	sum := sumup(intSlices...)

	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}
