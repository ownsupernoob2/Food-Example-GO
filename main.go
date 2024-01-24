package main

import (
	"fmt"
	"strings"
)

func main() {
	var foodType string
	var priceAdd float64
	menu := make(map[string]float64)
	total := 0.0

	fmt.Println("Enter menu items and prices (exit):")

	for {
		fmt.Print("Enter a food type (exit): ")
		fmt.Scan(&foodType)

		if strings.ToLower(foodType) == "exit" {
			break
		}

		fmt.Print("Enter the price for " + foodType + ": ")
		_, err := fmt.Scan(&priceAdd)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		menu[foodType] = priceAdd

		total += priceAdd
	}

	fmt.Println("Menu:")
	for foodType, price := range menu {
		fmt.Printf("%s: $%.2f\n", foodType, price)
	}

	fmt.Printf("Total: $%.2f\n", total)
}
