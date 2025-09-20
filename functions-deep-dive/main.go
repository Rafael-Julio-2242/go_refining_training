package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	doubledNumbers := transformNumbers(&numbers, double)
	tripledNumbers := transformNumbers(&numbers, triple)

	fmt.Println(doubledNumbers)
	fmt.Println(tripledNumbers)

	anonymousDoubled := func(number int) int {
		return number * 2
	}

	anonymousTripled := func(number int) int {
		return number * 3
	}

	anonymousDoubledNumbers := transformNumbers(&numbers, anonymousDoubled)
	anonymousTripledNumbers := transformNumbers(&numbers, anonymousTripled)

	fmt.Println("Anonymous Doubled: ", anonymousDoubledNumbers)
	fmt.Println("Anonymous Tripled: ", anonymousTripledNumbers)

	factoredDouble := createTransformer(2) // This is a closure
	factoredTriple := createTransformer(3)

	factoredDoubledNumbers := transformNumbers(&numbers, factoredDouble)
	factoredTripledNumbers := transformNumbers(&numbers, factoredTriple)

	fmt.Println("Factored Doubled: ", factoredDoubledNumbers)
	fmt.Println("Factored Tripled: ", factoredTripledNumbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, value := range *numbers {
		tNumbers = append(tNumbers, transform(value))
	}
	return tNumbers
}

func getTransformer() transformFn {
	return double
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
