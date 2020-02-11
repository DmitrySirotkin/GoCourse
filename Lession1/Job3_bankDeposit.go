package main

import "fmt"

func main() {
	countYears := 5
	count := 1

	fmt.Println("Вы запустили программу расчета вклада")

	sumDeposit := enterNum("Введите сумму вклада: ")
	bankPercent := enterNum("Введите процент, который будет начисляться ежегодно: ")

	fmt.Println("********************************************")
	fmt.Println("Поздравлем, вы открыли вклад на ", sumDeposit, " рублей, под ", bankPercent, "% годовых")
	fmt.Println("Ваши накопления составят:")

	for count <= countYears {
		sumDeposit = sumDeposit + (sumDeposit * (bankPercent / 100))
		fmt.Print("За ", count, "-й год ")
		fmt.Printf("%2.f руб.\n", sumDeposit)
		count++
	}
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
