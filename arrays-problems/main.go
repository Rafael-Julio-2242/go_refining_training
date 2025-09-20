package main

import "fmt"

func main() {

	hobbiesArray := task1()
	task2(hobbiesArray)
	hobbiesSlice := task3(hobbiesArray)
	task4(hobbiesSlice)
	goalsSlice := task5()
	task6(goalsSlice)
	task7()
}

func task1() [3]string {
	hobbiesArray := [3]string{
		"Playing Games",
		"Building Stuff",
		"Watching Animes / Series",
	}

	fmt.Println("(task 1) Hobbies: ", hobbiesArray)

	return hobbiesArray
}

func task2(hobbiesArray [3]string) {
	fmt.Println()
	fmt.Println("(Task 2) First element: ", hobbiesArray[0])
	fmt.Println("(Task 2) Second and Third Elements: ", hobbiesArray[1:])
	fmt.Println()
}

func task3(hobbiesArray [3]string) []string {
	hobbiesSlice := hobbiesArray[:2]
	secondHobbiesSlice := hobbiesArray[0:2]
	fmt.Println()
	fmt.Println("(Task 3) Hobbies Slice: ", hobbiesSlice)
	fmt.Println("(Task 3) Second Hobbies Slice: ", secondHobbiesSlice)
	return hobbiesSlice
}

func task4(hobbiesSlice []string) {
	hobbiesSlice = hobbiesSlice[1:]
	fmt.Println()
	fmt.Println("(Task 4) Hobbies Slice: ", hobbiesSlice)
}

func task5() []string {
	goalsSlice := []string{
		"Get better used to the language",
		"Be able to build stuff with Go",
	}

	fmt.Println("(Task 5) Goals Slice: ", goalsSlice)
	return goalsSlice
}

func task6(goalsSlice []string) {
	goalsSlice[1] = "Be able to build Awesome stuff with Go"
	fmt.Println()
	fmt.Println("(Task 6) Goals Slice: ", goalsSlice)
}

func task7() {
	productList := []Product{
		newProd(1, "Rope", 22.5),
		newProd(2, "Bottle", 5.6),
	}

	fmt.Println()
	fmt.Println("(Task 7) Products 'dynamic array': ", productList)

	productList = append(productList, newProd(3, "Controller", 35.6))
	fmt.Println("(Task 7) New Products 'dynamic array': ", productList)
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

type Product struct {
	id    int
	title string
	price float64
}

func newProd(id int, title string, price float64) Product {
	return Product{
		id,
		title,
		price,
	}
}
