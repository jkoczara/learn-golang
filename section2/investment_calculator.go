package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const FILE_NAME = "balance.txt"

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	for {
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawl amount:")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our bank")
			return
		}

		fmt.Println("Your choice:", choice)
	}
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(FILE_NAME)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balacne value.")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(FILE_NAME, []byte(balanceText), 0644)
}
