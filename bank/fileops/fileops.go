package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	// converts to string
	valueText := fmt.Sprint(value)
	// writes to file converting in byte array
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	// reads from file and puts in a byte array
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}
	// converts to string
	valueText := string(data)
	// converts to float
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}
	return value, nil
}
