package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var currentBalance float64
	var exit bool = false

	currentBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println("Error reading balance from file: ", err)
		panic("Can't continue...")
	}

	for {

		if exit {
			break
		}

		menu := generateMenu()
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

func generateMenu() (menu string) {
	menu = fmt.Sprintf(`
Welcome to Go Bank!
What do you want to do?
1. Check balance
2. Deposit Money
3. Withdraw Money
4. Exit
	`)

	return menu
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
	writeBalanceToFile(newBalance)
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
	writeBalanceToFile(newBalance)
	return newBalance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func readBalanceFromFile() (balance float64, err error) {
	bytes, err := os.ReadFile("balance.txt")

	if err != nil {
		return 0.0, err
	}

	balanceText := string(bytes)
	balance, err = strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0.0, err
	}

	return balance, nil
}
