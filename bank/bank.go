package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic("Exiting")
	}

	fmt.Printf("Welcome to Go Bank %s!\n", randomdata.SillyName())
	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}
			accountBalance += deposit
			fmt.Println("Your updated balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		case 3:
			fmt.Print("Choose how much you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than what you have.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your updated balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, balanceFile)
		default:
			fmt.Println("Goodbye!")
			return
		}

		fmt.Println("Thanks for choosing our bank")
	}

}
