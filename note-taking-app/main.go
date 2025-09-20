package main

import (
	"bufio"
	"fmt"
	"go-note-taking-app/note"
	"go-note-taking-app/todo"
	"os"
	"strings"
)

type saver interface {
	SaveFile() error
}

type outputtable interface {
	saver
	Display()
}

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

	todoText, todoErr := getUserInput("Todo Text: ")

	userNote, noteErr := note.New(title, content)

	if noteErr != nil {
		panic(noteErr)
	}

	if todoErr != nil {
		panic(todoErr)
	}

	userTodo := todo.New(todoText)

	userNoteError := outputData(userNote)
	userTodoError := outputData(userTodo)

	if userNoteError != nil {
		panic(userNoteError)
	}

	if userTodoError != nil {
		panic(userTodoError)
	}

	fmt.Println("Note Saved!")
}

func outputData(data outputtable) error {
	data.Display()
	return data.SaveFile()
}

func Save(sv saver) error {
	err := sv.SaveFile()

	if err != nil {
		return err
	}

	return nil
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
