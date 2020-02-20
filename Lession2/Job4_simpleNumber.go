package main

import "fmt"

func main() {
	var simpleNum []int
	countArr := enterNum("Введите максимальное число для нахождения простых чисел: ")
	allNum := arrayBuilding(countArr)

	for i := 0; i <= len(allNum)-1; i++ {
		if allNum[i] > 1 {
			simpleNum = append(simpleNum, allNum[i])
			for j := i; j <= len(allNum)-1; j = j + i {
				allNum[j] = 0
			}
		}
	}

	fmt.Println(simpleNum)

}

func arrayBuilding(count int) []int {
	var allNum []int
	for i := 0; i <= count; i++ {
		allNum = append(allNum, i)
	}
	return allNum
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
