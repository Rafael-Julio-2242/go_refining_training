package main

import (
	"bufio"
	"fmt"
	"go-note-taking-app/note"
	"os"
	"strings"
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

	userNote, noteErr := note.New(title, content)

	if noteErr != nil {
		panic(noteErr)
	}

	userNote.ShowNoteInfo()

	fmt.Println("Saving note...")

	err := userNote.SaveFile()

	if err != nil {
		panic("Error saving note: " + err.Error())
	}

	fmt.Println("Note Saved!")

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	value, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value, nil
}
