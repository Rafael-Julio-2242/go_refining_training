package main

import (
	"fmt"
	"go-simple-bank/fileops"
)

func main() {

	var currentBalance float64
	var exit bool = false

	currentBalance, err := fileops.ReadFloatFromFile("balance.txt")

	if err != nil {
		fmt.Println("Error reading balance from file: ", err)
		panic("Can't continue...")
	}

	for {

		if exit {
			break
		}

		menu := GenerateMenu()
		fmt.Println(menu)

		var action int
		fmt.Print("Your choice: ")
		fmt.Scan(&action)

		switch action {
		case 1:
			showBalance(currentBalance)
		case 2:
			currentBalance = depositMoney(currentBalance)
		case 3:
			currentBalance = withdrawMoney(currentBalance)
		case 4:
			fmt.Println("Exiting...")
			exit = true
		default:
			fmt.Println("Invalid Action!")
		}

	}

}

func showBalance(balance float64) {
	fmt.Println("Your balance is ", balance)
}

func depositMoney(currentBalance float64) (newBalance float64) {
	var amount float64
	fmt.Print("Enter the amount to deposit: ")
	fmt.Scan(&amount)

	newBalance = currentBalance + amount

	fmt.Println("Depoisted ", amount)
	fmt.Println("New Balance is ", newBalance)
	fileops.WriteFloatToFile("balance.txt", newBalance)
	return newBalance
}

func withdrawMoney(currentBalance float64) (newBalance float64) {
	var amount float64
	fmt.Print("Enter the amount to Withdraw: ")
	fmt.Scan(&amount)

	newBalance = currentBalance

	if amount > currentBalance {
		fmt.Println("Cannot Withdraw more than the current balance!")
		return newBalance
	}

	newBalance -= amount
	fmt.Printf("Withdrawed %v\n", amount)
	fmt.Printf("New Balance is %v\n", newBalance)
	fileops.WriteFloatToFile("balance.txt", newBalance)
	return newBalance
}
