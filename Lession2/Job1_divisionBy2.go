package main

import (
	"fmt"
)

func main() {
	calcNum := enterNum("Введите число для определения четности: ")
	if calcNum%2 == 0 {
		fmt.Println("Вы ввели четное число!")
	} else {
		fmt.Println("Вы ввели не четное число!")
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
