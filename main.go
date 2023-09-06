package main

import (
	"fmt"

	"github.com/matheusgb/go-jobs-crawler/scrapers"
)

func main() {
	var option int
	fmt.Println("Select an option: ")
	fmt.Println("1 - Linkedin")
	fmt.Scanln(&option)

	switch option {
	case 1:
		scrapers.LinkedinScrap()
	default:
		fmt.Println("Invalid option")
	}
}
