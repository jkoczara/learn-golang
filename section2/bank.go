package main

import (
	"fmt"

	"exmaple.com/bank/fileops"
)

const FILE_NAME = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(FILE_NAME)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		//panic(err)
	}

	presentOptions()

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
			fileops.WriteBalanceToFile(accountBalance, FILE_NAME)
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
			fileops.WriteBalanceToFile(accountBalance, FILE_NAME)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our bank")
			return
		}

		fmt.Println("Your choice:", choice)
	}
}
