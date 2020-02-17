package main

import (
	"fmt"
)

func main() {
	calcNum := enterNum("Введите число для определения делится ли число без остатка на 3: ")
	if calcNum%3 == 0 {
		fmt.Println("Число делится на 3 без остатка!")
	} else {
		fmt.Println("Число не делится на 3 без остатка!")
	}
}

func enterNum(invitToEnter string) int {
	var num int
	fmt.Print(invitToEnter)
	for {
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Print("Пожалуйста, ведите число: ")
		} else {
			return num
		}
	}
}
