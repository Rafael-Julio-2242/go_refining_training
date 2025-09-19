package main

func GenerateMenu() (menu string) {
	menu = `
Welcome to Go Bank!
What do you want to do?
1. Check balance
2. Deposit Money
3. Withdraw Money
4. Exit
	`

	return menu
}
