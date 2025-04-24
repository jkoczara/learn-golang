package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fielName string) (float64, error) {
	data, err := os.ReadFile(fielName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	value, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func WriteBalanceToFile(value float64, fielName string) {
	balanceText := fmt.Sprint(value)
	os.WriteFile(fielName, []byte(balanceText), 0644)
}
