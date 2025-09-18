package main

import (
	"errors"
	"fmt"
	"os"
)

const outputFileName = "calculations.txt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	revenue, err = getUserInput("Enter Revenue: ")
	if err != nil {
		panic(err)
	}

	expenses, err = getUserInput("Enter expenses: ")
	if err != nil {
		panic(err)
	}

	taxRate, err = getUserInput("Enter Tax Rate: ")
	if err != nil {
		panic(err)
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fullMessage := createOutputMessage(ebt, profit, ratio)

	writeCalculationsInFile(fullMessage)

	fmt.Println(fullMessage)
}

func getUserInput(message string) (userInput float64, err error) {
	fmt.Print(message)
	fmt.Scan(&userInput)

	if userInput < 0 {
		err := errors.New("input must be a positive number")
		return 0, err
	}

	if userInput == 0 {
		err := errors.New("input cannot be 0")
		return 0, err
	}

	return userInput, nil
}

func calculate(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func writeCalculationsInFile(output string) {
	outputBytes := []byte(output)
	os.WriteFile(outputFileName, outputBytes, 0644)
}

func createOutputMessage(ebt float64, profit float64, ratio float64) (message string) {
	message = fmt.Sprintf(`
EBT: %.2f
Profit: %.2f
Ratio: %.2f
	`, ebt, profit, ratio)

	return message
}
