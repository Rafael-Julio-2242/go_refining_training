package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5

	var investmentAmount float64
	var years int
	var returnRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter return rate: ")
	fmt.Scan(&returnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	finalAmount := investmentAmount * math.Pow(1+(returnRate/100), float64(years))
	futureRealValue := finalAmount / math.Pow(1+(inflationRate/100), float64(years))

	fmt.Printf("Final amount: %.2f\n", finalAmount)
	fmt.Printf("Future real value: %.2f\n", futureRealValue)

}
