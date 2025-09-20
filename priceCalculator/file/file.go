package file

import (
	"encoding/json"
	"errors"
	"go-price-calculator/tax"
	"os"
	"strconv"
	"strings"
)

const permissionNumber = 0644

func GetFormatedDataFromFile(fileName string) ([]float64, error) {
	var data []float64

	fileBytes, readFileErr := os.ReadFile(fileName)

	if readFileErr != nil {
		return nil, errors.New("Error Reading File: " + readFileErr.Error())
	}

	fileString := string(fileBytes)

	fileStringValues := strings.Split(fileString, "\n")

	for _, value := range fileStringValues {
		if value == "" {
			continue
		}
		floatValue, convErr := strconv.ParseFloat(value, 64)

		if convErr != nil {
			return nil, errors.New("Error converting values: " + convErr.Error())
		}

		data = append(data, floatValue)
	}

	return data, nil
}

func WriteResultJsonFile(fileName string, results []tax.TaxResult) error {
	jsonBytes, err := json.Marshal(results)

	if err != nil {
		return errors.New("Error on Transforming the results in JSON bytes: " + err.Error())
	}

	writeErr := os.WriteFile(fileName, jsonBytes, permissionNumber)

	if writeErr != nil {
		return errors.New("Error on Creating the Results JSON: " + writeErr.Error())
	}

	return nil
}
