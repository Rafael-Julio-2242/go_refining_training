package main

import "fmt"

func main() {

	factorOfFive := factorial(5)

	fmt.Println(factorOfFive)

}

func factorial(number int) int {
	if number > 1 {
		return number * factorial(number-1)
	} else {
		return number
	}
}
