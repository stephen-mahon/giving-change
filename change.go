package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Number of coins returned from denominations of 1, 5, 10, 25, 100, and 500 units")
		fmt.Println("Usage: change <Total ammount>")
		fmt.Println("Example: change 12")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter one arguments! Type /help for help.")
		} else {
			ammount, _ := strconv.ParseInt(args[0], 0, 64)
			fmt.Println(coinsReturned(int(ammount)))
		}
	}
}

func coinsReturned(ammount int) int {
	coins := [6]int{500, 100, 25, 10, 5, 1}
	change := ammount
	var coinNum int = 0
	for _, v := range coins {
		for change >= v {
			change -= v
			coinNum++
		}
	}
	return coinNum
}
