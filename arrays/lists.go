package main

import "fmt"

func main() {
	var productNames [4]string
	prices := [4]float64{10.99, 9.99, 40.0, 20.0}
	fmt.Println(prices)
	fmt.Println(productNames)

	productNames[2] = "A Carpet"

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlitedPrices := featuredPrices[:1]
	fmt.Println(highlitedPrices)
	fmt.Println(prices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

	highlitedPrices = highlitedPrices[:3]
	fmt.Println(highlitedPrices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

}
