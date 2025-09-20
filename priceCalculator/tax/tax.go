package tax

type TaxResult struct {
	Tax          float64   `json:"tax"`
	ResultPrices []float64 `json:"result_prices"`
}

func NewTaxResult(taxValue float64, resultPrices []float64) TaxResult {
	return TaxResult{
		Tax:          taxValue,
		ResultPrices: resultPrices,
	}
}

func GetTaxValues() []float64 {
	return []float64{0.0, 10.0, 20.0}
}
