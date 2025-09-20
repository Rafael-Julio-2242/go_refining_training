package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

const permissionNumber = 0644

type Todo struct {
	Text string `json:"text"`
}

func New(text string) *Todo {
	return &Todo{
		Text: text,
	}
}

func (t *Todo) ShowNoteInfo() {
	fmt.Println(t.Text)
}

func (t Todo) ToJson() ([]byte, error) {
	jsonValue, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return jsonValue, nil
}

func (t *Todo) SaveFile() error {

	jsonValue, jsonErr := t.ToJson()

	if jsonErr != nil {
		return jsonErr
	}

	fileName := "todo.json"

	return os.WriteFile(fileName, jsonValue, permissionNumber)
}
