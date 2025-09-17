package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Enter Revenue: ")

	expenses = getUserInput("Enter expenses: ")

	taxRate = getUserInput("Enter Tax Rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fullMessage := createOutputMessage(ebt, profit, ratio)

	fmt.Println(fullMessage)
}

func getUserInput(message string) (userInput float64) {
	fmt.Print(message)
	fmt.Scan(&userInput)
	return userInput
}

func calculate(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func createOutputMessage(ebt float64, profit float64, ratio float64) (message string) {
	message = fmt.Sprintf(`
EBT: %.2f
Profit: %.2f
Ratio: %.2f
	`, ebt, profit, ratio)

	return message
}
