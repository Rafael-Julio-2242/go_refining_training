package main

import (
	"fmt"
	"structs-examples/user"
)

func main() {

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your Birth Date (DD/MM/YYYY): ")

	appUser, err := user.New(firstName, lastName, birthdate)

	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin("test@example.com", "123456")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(message string) (value string) {
	fmt.Print(message)
	fmt.Scanln(&value)
	return value
}
