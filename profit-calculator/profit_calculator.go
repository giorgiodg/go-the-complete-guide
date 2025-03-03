package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenues, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting")
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting")
		return
	}
	taxRate, err := getUserInput("Tax rate: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenues, expenses, taxRate)
	storeResultsToFile(ebt, profit, ratio)

	fmt.Printf("EBT %1.f\n", ebt)
	fmt.Printf("Profit %1.f\n", profit)
	fmt.Printf("Ratio %1.f\n", ratio)

}

func getUserInput(inputLabel string) (float64, error) {
	var inputVar float64
	fmt.Print(inputLabel)
	fmt.Scan(&inputVar)
	if inputVar <= 0 {
		return 0, errors.New("value must be greater than 0")
	}
	return inputVar, nil
}

func calculateFinancials(revenues, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenues - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResultsToFile(ebt, profit, ratio float64) {
	inputString := fmt.Sprintf("EBT: %1.f\nProfit: %1.f\nRatio %3.f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(inputString), 0644)
}
