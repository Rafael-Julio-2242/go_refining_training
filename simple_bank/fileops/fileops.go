package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func ReadFloatFromFile(fileName string) (value float64, err error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		customError := errors.New("error to read value from file")
		return 0.0, customError
	}

	valueText := string(bytes)
	value, err = strconv.ParseFloat(valueText, 64)

	if err != nil {
		customError := errors.New("error to write value in file")
		return 0.0, customError
	}

	return value, nil
}
