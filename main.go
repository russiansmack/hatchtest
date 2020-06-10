package main

import (
	"fmt"

	tcaw "github.com/russiansmack/thecatapiwrapper"
)

func main() {

	fmt.Println("Hello, Hatch!")

	c := tcaw.NewClient("7a1768b0-1600-4c55-9769-83721284ab92")

	fmt.Println("--- GET /images/search ---")

	options := tcaw.ImageSearchOptions{}
	options.Limit = 5
	options.Order = "ASC"
	options.Page = 0

	images, err := c.GetImageSearch(options)
	if err != nil {
		fmt.Println("badjuju")
	}
	fmt.Printf("%+v\n", images)

	fmt.Println("--- GET /categories ---")

	categories, err := c.GetCategories()
	if err != nil {
		fmt.Println("badjuju")
	}
	fmt.Printf("%+v\n", categories)

}
