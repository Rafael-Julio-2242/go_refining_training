package main

import "fmt"

func main() {

	age := 32

	fmt.Println("Age: ", age)

	fmt.Println("Adult Years: ", getAdultYears(age))

}

func getAdultYears(age int) int {
	return age - 18
}
