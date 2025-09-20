package main

import (
	"go-price-calculator/file"
	"go-price-calculator/tax"
)

const fileName = "prices.txt"

func main() {
	var results []tax.TaxResult

	prices, loadPricesErr := file.GetFormatedDataFromFile(fileName)

	if loadPricesErr != nil {
		panic(loadPricesErr)
	}

	taxRates := tax.GetTaxValues()

	for _, taxValue := range taxRates {
		var taxPrices []float64

		for _, price := range prices {
			taxPrices = append(taxPrices, calculateTaxedPrice(price, taxValue))
		}

		res := tax.NewTaxResult(taxValue, taxPrices)

		results = append(results, res)
	}

	file.WriteResultJsonFile("results.json", results)
}

func calculateTaxedPrice(price float64, tax float64) float64 {
	return price + (price * (tax / 100))
}
