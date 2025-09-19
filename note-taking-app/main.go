package main

import (
	"errors"
	"fmt"
)

func main() {

	title, titleErr := getUserInput("Note title: ")

	if titleErr != nil {
		fmt.Println("Title error: ", titleErr)
		return
	}

	content, contentErr := getUserInput("Note content: ")

	if contentErr != nil {
		fmt.Println("Content error: ", contentErr)
	}

	fmt.Println(title, content)

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}
