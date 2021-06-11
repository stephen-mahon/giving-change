package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Change given from denominations of 1, 5, 10, 25, 100, and 500 units")
		fmt.Println("Usage: change-given <Total ammount>")
		fmt.Println("Example: change 12")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one arguments! Type /help for help.")
		} else {
			ammount, _ := strconv.ParseFloat(args[0], 64)
			fmt.Println(changeGiven(ammount))
		}
	}
}

func changeGiven(ammount float64) float64 {
	return ammount
}
