package main

import "fmt"

func main() {
	// 1
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[0:2]
	// mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	goals := []string{"upskilling", "career pivot"}
	fmt.Println(goals)

	goals[1] = "career growth"
	fmt.Println(goals)

	// 6
	goals = append(goals, "open source giveback")
	fmt.Println(goals)

	// 7
	productList := []Product{
		{"Product A", 1, 0.1},
		{"Product B", 2, 100.0}}
	fmt.Println(productList)

	// newProduct := Product{"Product C", 3, 0.5}
	// productList = append(productList, newProduct)
	productList = append(productList, Product{"Product C", 3, 0.5})
	otherProductList := []Product{{"Product D", 4, 0.5}, {"Product E", 5, 10.5}}
	fmt.Println(productList)
	fmt.Println(otherProductList)
	productList = append(productList, otherProductList...)
	fmt.Println(productList)
}

type Product struct {
	title string
	id    int
	price float64
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
