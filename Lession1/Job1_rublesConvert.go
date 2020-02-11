package main

import "fmt"

func main() {
	var dollarRate float64 = 64.45

	fmt.Println("Вы запустили конвектор валюты, стоимость 1 $ равна", dollarRate, "₽")
	rubleTotal := enterNum("Введите сумму для конвертации в рублях: ")
	fmt.Printf("Итог конвертации: %.2f $", rubleTotal/dollarRate)
}

func enterNum(invitToEnter string) float64 {
	var num float64 = -1
	fmt.Print(invitToEnter)
	for num < 0 {
		if _, err := fmt.Scan(&num); err != nil || num <= 0 {
			fmt.Print("Пожалуйста, положительное число: ")
			num = -1
		}
	}
	return num
}
