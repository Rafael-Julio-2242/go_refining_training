package main

import "fmt"

type floatMap map[string]float64

func main() {

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Linkedin"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "Amazon Web Services")

	fmt.Println(websites)

	fmt.Println("--------------------------")

	courseRatings := make(map[string]float64, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	fmt.Println(courseRatings)
}
