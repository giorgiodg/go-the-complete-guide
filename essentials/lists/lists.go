package lists

import "fmt"

func main() {

	var productNames [4]string = [4]string{"A book"}

	prices := [4]float64{1.1, 10.99, 9.99, 44.4}
	fmt.Println(prices)

	productNames[3] = "A Carpet"
	fmt.Println(prices[2])
	fmt.Println(productNames)

	featuredPrices := prices[1:]

	fmt.Println(featuredPrices)
	fmt.Println(len(featuredPrices))
	fmt.Println(cap(featuredPrices))
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	updatedPrices := append(prices, 10.99)

	fmt.Println(updatedPrices, prices)
}
