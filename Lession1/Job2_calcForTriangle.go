package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вы запустили программу расчета параметров прямоугольного треугольника")
	longTriangle1 := enterNum("Введите длину 1-го катета: ")
	longTriangle2 := enterNum("Введите длину 2-го катета: ")
	fmt.Println("--------------------------------")
	fmt.Println("Площадь треугольника: ", longTriangle1*longTriangle2/2)
	fmt.Printf("Периметр треугольника: %.2f\n", calcPerim(longTriangle1, longTriangle2))
	fmt.Printf("Гипотенуза треугольника: %.2f", calcGip(longTriangle1, longTriangle2))
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

func calcPerim(long1 float64, long2 float64) float64 {
	return long1 + long2 + calcGip(long1, long2)
}

func calcGip(long1 float64, long2 float64) float64 {
	return math.Sqrt(math.Pow(long1, 2) + math.Pow(long2, 2))
}
