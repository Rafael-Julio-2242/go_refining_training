package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const noteFileName = "note.json"
const permissionNumber = 0644

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("note must have title and content")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) ShowNoteInfo() {
	fmt.Printf("\n\nTitle: %v\nContent: %v\n\n", n.Title, n.Content)
}

func (n Note) ToJson() ([]byte, error) {
	jsonValue, err := json.Marshal(n)

	if err != nil {
		return nil, err
	}

	return jsonValue, nil
}

func (n *Note) SaveFile() error {

	jsonValue, jsonErr := n.ToJson()

	if jsonErr != nil {
		return jsonErr
	}

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + "_" + noteFileName

	return os.WriteFile(fileName, jsonValue, permissionNumber)
}
