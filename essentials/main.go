package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	text, sum := sumup("a", 1, 10, 15, 40)

	anotherText, anotherSum := sumup("b", numbers...)
	fmt.Println(text, sum)
	fmt.Println(anotherText, anotherSum)
}

func sumup(text string, numbers ...int) (string, int) {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return text, sum
}
